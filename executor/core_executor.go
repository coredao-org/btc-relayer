package executor

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	cgccaller "github.com/coredao-org/btc-relayer/executor/cc"
	"github.com/coredao-org/btcpowermirror/lightmirror"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"

	brcommon "github.com/coredao-org/btc-relayer/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/jinzhu/gorm"

	relayercommon "github.com/coredao-org/btc-relayer/common"
	config "github.com/coredao-org/btc-relayer/config"
	"github.com/coredao-org/btc-relayer/executor/relayerhub"
)

type COREClient struct {
	COREClient    *ethclient.Client
	Provider      string
	CurrentHeight int64
	UpdatedAt     time.Time
}

type COREExecutor struct {
	mutex       sync.RWMutex
	db          *gorm.DB
	btcExecutor *BTCExecutor
	clientIdx   int
	coreClients []*COREClient
	privateKey  *ecdsa.PrivateKey
	txSender    common.Address
	cfg         *config.Config
}

func getPrivateKey(cfg *config.COREConfig) (*ecdsa.PrivateKey, error) {
	var privateKey string

	privateKey = cfg.PrivateKey

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func initClients(providers []string) []*COREClient {
	clients := make([]*COREClient, 0)

	for _, provider := range providers {
		client, err := ethclient.Dial(provider)
		if err != nil {
			panic("new eth client error")
		}
		clients = append(clients, &COREClient{
			COREClient: client,
			Provider:   provider,
			UpdatedAt:  time.Now(),
		})
	}

	return clients
}

func NewCOREExecutor(cfg *config.Config) (*COREExecutor, error) {
	privKey, err := getPrivateKey(&cfg.COREConfig)
	if err != nil {
		return nil, err
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("get public key error")
	}
	txSender := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &COREExecutor{
		db:          nil,
		btcExecutor: nil,
		clientIdx:   0,
		coreClients: initClients(cfg.COREConfig.Providers),
		privateKey:  privKey,
		txSender:    txSender,
		cfg:         cfg,
	}, nil
}

func (executor *COREExecutor) GetClient() *ethclient.Client {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.coreClients[executor.clientIdx].COREClient
}

func (executor *COREExecutor) SwitchCOREClient() {
	executor.mutex.Lock()
	defer executor.mutex.Unlock()
	executor.clientIdx++
	if executor.clientIdx >= len(executor.coreClients) {
		executor.clientIdx = 0
	}
	relayercommon.Logger.Infof("Switch to provider: %s", executor.cfg.COREConfig.Providers[executor.clientIdx])
}

func (executor *COREExecutor) GetLatestBlockHeight(client *ethclient.Client) (int64, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	block, err := client.BlockByNumber(ctxWithTimeout, nil)
	if err != nil {
		return 0, err
	}
	return block.Number().Int64(), nil
}

func (executor *COREExecutor) UpdateClients() {
	for {
		for _, client := range executor.coreClients {
			if time.Since(client.UpdatedAt).Seconds() > executor.cfg.COREConfig.DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessible", client.Provider)
				relayercommon.Logger.Error(msg)
				config.SendTelegramMessage(executor.cfg.AlertConfig.Identity, executor.cfg.AlertConfig.TelegramBotId, executor.cfg.AlertConfig.TelegramChatId, msg)
			}
			height, err := executor.GetLatestBlockHeight(client.COREClient)
			if err != nil {
				relayercommon.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			client.CurrentHeight = height
			client.UpdatedAt = time.Now()
		}
		//relayercommon.Logger.Infof("Start to monitor core data-seeds health")

		highestHeight := int64(0)
		highestIdx := 0
		for idx := 0; idx < len(executor.coreClients); idx++ {
			if executor.coreClients[idx].CurrentHeight > highestHeight {
				highestHeight = executor.coreClients[idx].CurrentHeight
				highestIdx = idx
			}
		}

		if executor.coreClients[executor.clientIdx].CurrentHeight+FallBehindThreshold < highestHeight {
			executor.mutex.Lock()
			executor.clientIdx = highestIdx
			executor.mutex.Unlock()
		}
		time.Sleep(time.Duration(executor.cfg.COREConfig.SleepSecond) * time.Second)
	}
}

func (executor *COREExecutor) getTransactor(nonce uint64) (*bind.TransactOpts, error) {
	chainId, err := executor.GetClient().ChainID(context.Background())
	if err != nil {
		return nil, nil
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(executor.privateKey, chainId)
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = executor.cfg.COREConfig.GasLimit
	if executor.cfg.COREConfig.GasPrice == 0 {
		txOpts.GasPrice = big.NewInt(DefaultGasPrice)
	} else {
		txOpts.GasPrice = big.NewInt(int64(executor.cfg.COREConfig.GasPrice))
	}
	return txOpts, nil
}

func (executor *COREExecutor) GetChainTip() (*chainhash.Hash, error) {
	callOpts, err := executor.getCallOpts()
	instance, err := cgccaller.NewCGCCaller(pcsAddr, executor.GetClient())
	chainTip, err := instance.GetChainTip(callOpts)

	return chainTip, err
}

func (executor *COREExecutor) getCallOpts() (*bind.CallOpts, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}
	return callOpts, nil
}

/**
sync BTCLightMirror
*/
func (executor *COREExecutor) SyncBTCLightMirror(task *relayercommon.Task) (common.Hash, error) {
	nonce, err := executor.GetClient().PendingNonceAt(context.Background(), executor.txSender)
	if err != nil {
		return common.Hash{}, err
	}
	txOpts, err := executor.getTransactor(nonce)
	if err != nil && txOpts != nil {
		return common.Hash{}, err
	}

	mirror := NewBtcLightMirror(task.BLOCK)

	for {
		txHash, err := executor.syncBtcHeader(mirror, task.BlockHash)
		if err != nil {
			return common.Hash{}, err
		}

		brcommon.Logger.Infof("submit transaction, blockHash:" + task.BlockHash.String() + " height:" + Int64ToString(task.Height) + ",txHash:" + txHash.String() + ", start to check relaying result")

		relayed, retry, err := executor.CheckSuccessRelayed(task.BlockHash, txHash)

		if relayed || !retry {
			return common.Hash{}, err
		}

		if retry {
			executor.IncreaseGas()
		}

	}

	return common.Hash{}, err
}

/**
increase gas
*/
func (executor *COREExecutor) IncreaseGas() {
	executor.cfg.COREConfig.GasLimit += executor.cfg.COREConfig.GasIncrease
	brcommon.Logger.Infof("gas not enough, increase gas to:" + strconv.FormatUint(executor.cfg.COREConfig.GasLimit, 10))
}

/**
check if btc block is successfully relayed
return bool:relayed success bool:retry
*/
func (executor *COREExecutor) CheckSuccessRelayed(btcBlockHash *chainhash.Hash, coreTxHash common.Hash) (bool, bool, error) {

	for {
		//CheckBlockRelayed
		relayed, err := executor.CheckBlockRelayed(btcBlockHash)
		if err == nil && relayed {
			submitter, err := executor.QuerySubmitters(btcBlockHash)
			if err == nil && submitter != "" {
				brcommon.Logger.Infof("successful, relayed by:[" + submitter + "]")
			} else {
				brcommon.Logger.Infof("successful")
			}
			return true, false, nil
		}

		//Check TX
		txRecipient, err := executor.GetTxRecipient(coreTxHash)

		//failed, get revert reason
		if err == nil {
			if txRecipient.Status == 0 {
				tx, _, err := executor.TransactionByHash(coreTxHash)
				if err == nil {

					//out of gas
					if tx.Gas() == txRecipient.GasUsed {
						brcommon.Logger.Infof("out of gas, retry")
						return false, true, nil
					}
					brcommon.Logger.Infof("failed")
					return false, false, fmt.Errorf("tx failed")
				}
			}
		}

		relayercommon.Logger.Infof("relaying, continue to check")
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

}

func serializeBtcLightMirror(mirror *lightmirror.BtcLightMirrorV2) ([]byte, error) {
	var b bytes.Buffer
	mirror.Serialize(&b)
	bb := b.Bytes()

	return bb, nil
}

func (executor *COREExecutor) IsRelayer() (bool, error) {
	instance, err := relayerhub.NewRelayerhub(relayerHubContractAddr, executor.GetClient())
	if err != nil {
		return false, err
	}

	callOpts, err := executor.getCallOpts()
	if err != nil {
		return false, err
	}

	isRelayer, err := instance.IsRelayer(callOpts, executor.txSender)
	if err != nil {
		return false, err
	}
	return isRelayer, nil
}

func (executor *COREExecutor) RegisterRelayer() (common.Hash, error) {
	nonce, err := executor.GetClient().PendingNonceAt(context.Background(), executor.txSender)
	if err != nil {
		return common.Hash{}, err
	}
	txOpts, err := executor.getTransactor(nonce)
	if err != nil {
		return common.Hash{}, err
	}

	instance, err := relayerhub.NewRelayerhub(relayerHubContractAddr, executor.GetClient())
	if err != nil {
		return common.Hash{}, err
	}

	txOpts.Value = big.NewInt(1).Mul(big.NewInt(100), big.NewInt(1e18)) //100 Core
	tx, err := instance.Register(txOpts)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (executor *COREExecutor) EthCall(tx *types.Transaction, blockNumber *big.Int) ([]byte, error) {
	msg := ethereum.CallMsg{
		From:     executor.txSender,
		To:       tx.To(),
		GasPrice: tx.GasPrice(),
		Gas:      tx.Gas(),
		//GasTipCap: tx.GasTipCap(),
		//GasFeeCap: tx.GasFeeCap(),
		Value: tx.Value(),
		Data:  tx.Data(),
	}
	return executor.GetClient().CallContract(context.Background(), msg, blockNumber)
}

func (executor *COREExecutor) TransactionByHash(txHash common.Hash) (*types.Transaction, bool, error) {
	return executor.GetClient().TransactionByHash(context.Background(), txHash)
}

func (executor *COREExecutor) GetTxRecipient(txHash common.Hash) (*types.Receipt, error) {
	return executor.GetClient().TransactionReceipt(context.Background(), txHash)
}

func (executor *COREExecutor) GetRelayerBalance() (*big.Int, error) {
	return executor.GetClient().BalanceAt(context.Background(), executor.txSender, nil)
}

func (executor *COREExecutor) CheckBlockRelayed(blockHash *chainhash.Hash) (bool, error) {

	bHash := RevertHash(blockHash)

	callOpts, err := executor.getCallOpts()
	instance, err := cgccaller.NewCGCCaller(pcsAddr, executor.GetClient())
	result, err := instance.IsHeaderSynced(callOpts, bHash)
	//_, err := executor.CallContext(executor.GetClient(), result, context.Background(), "isHeaderSynced", blockHash)

	return result, err
}

func (executor *COREExecutor) QuerySubmitters(blockHash *chainhash.Hash) (string, error) {

	bHash := RevertHash(blockHash)

	callOpts, err := executor.getCallOpts()
	instance, err := cgccaller.NewCGCCaller(pcsAddr, executor.GetClient())
	result, err := instance.QuerySubmitters(callOpts, bHash)

	return result, err
}

func (executor *COREExecutor) syncBtcHeader(btcLightMirror *lightmirror.BtcLightMirrorV2, blockHash *chainhash.Hash) (common.Hash, error) {
	nonce, err := executor.GetClient().PendingNonceAt(context.Background(), executor.txSender)
	callOpts, err := executor.getTransactor(nonce)

	instance, err := cgccaller.NewCGCCaller(pcsAddr, executor.GetClient())

	bts, err := serializeBtcLightMirror(btcLightMirror)
	txHash, err := instance.SyncBtcHeader(callOpts, bts)

	if err != nil {
		log.Println("sync btc header failed, hash:" + blockHash.String())
	}

	return txHash, err
}

// callContext
func (executor *COREExecutor) CallContext(ec *ethclient.Client, result interface{}, ctx context.Context, method string, args ...interface{}) (interface{}, error) {
	client := relayercommon.ReflectField(ec, "c").Elem().Interface().(*rpc.Client)
	err := client.CallContext(ctx, &result, method, args)
	if err != nil {
		return nil, err
	}
	return result, err
}
