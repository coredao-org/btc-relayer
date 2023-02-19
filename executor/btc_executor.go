package executor

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/coredao-org/btc-relayer/common"
	config "github.com/coredao-org/btc-relayer/config"
)

type BTCClient struct {
	BTCClient     *rpcclient.Client
	Provider      string
	CurrentHeight int64
	UpdatedAt     time.Time
}

type BTCExecutor struct {
	mutex         sync.RWMutex
	clientIdx     int
	HighestHeight int64
	BTCClients    []*BTCClient
	Config        *config.Config
}

func initBTCClients(providers []config.BTCRpcAddrs) []*BTCClient {
	btcClients := make([]*BTCClient, 0)
	for _, provider := range providers {

		// create new client instance
		btcClient, err := rpcclient.New(&rpcclient.ConnConfig{
			HTTPPostMode: true,
			DisableTLS:   true,
			Host:         provider.Host,
			User:         provider.User,
			Pass:         provider.Pass,
		}, nil)
		if err != nil {
			log.Fatalf("error creating new btc client: %v", err)
		}

		btcClients = append(btcClients, &BTCClient{
			BTCClient: btcClient,
			Provider:  provider.Host,
			UpdatedAt: time.Now(),
		})
	}
	return btcClients
}

func NewBTCExecutor(cfg *config.Config) (*BTCExecutor, error) {
	return &BTCExecutor{
		clientIdx:  0,
		BTCClients: initBTCClients(cfg.BTCConfig.RpcAddrs),
		Config:     cfg,
	}, nil
}

func (executor *BTCExecutor) GetClient() *rpcclient.Client {
	executor.mutex.RLock()
	defer executor.mutex.RUnlock()
	return executor.BTCClients[executor.clientIdx].BTCClient
}

func (executor *BTCExecutor) SwitchBTClient() {
	executor.mutex.Lock()
	defer executor.mutex.Unlock()
	executor.clientIdx++
	if executor.clientIdx >= len(executor.BTCClients) {
		executor.clientIdx = 0
	}
	common.Logger.Infof("Switch to RPC endpoint: %s", executor.Config.BTCConfig.RpcAddrs[executor.clientIdx])
}

func (executor *BTCExecutor) GetLatestBlockHeight(client *rpcclient.Client) (int64, error) {
	height, err := client.GetBlockCount()
	if err != nil {
		return 0, err
	}
	return height, nil
}

func (executor *BTCExecutor) GetBlockHash(client *rpcclient.Client, height int64) (*chainhash.Hash, error) {
	hash, err := client.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (executor *BTCExecutor) GetBlock(client *rpcclient.Client, hash *chainhash.Hash) (*wire.MsgBlock, error) {
	block, err := client.GetBlock(hash)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (executor *BTCExecutor) UpdateClients() {
	for {
		for _, btcClient := range executor.BTCClients {
			if time.Since(btcClient.UpdatedAt).Seconds() > executor.Config.BTCConfig.DataSeedDenyServiceThreshold {
				msg := fmt.Sprintf("data seed %s is not accessible", btcClient.Provider)
				common.Logger.Error(msg)
				config.SendTelegramMessage(executor.Config.AlertConfig.Identity, executor.Config.AlertConfig.TelegramBotId, executor.Config.AlertConfig.TelegramChatId, msg)
			}
			height, err := executor.GetLatestBlockHeight(btcClient.BTCClient)
			if err != nil {
				common.Logger.Errorf("get latest block height error, err=%s", err.Error())
				continue
			}
			btcClient.CurrentHeight = height
			btcClient.UpdatedAt = time.Now()
		}
		highestHeight := int64(0)
		highestIdx := 0
		for idx := 0; idx < len(executor.BTCClients); idx++ {
			if executor.BTCClients[idx].CurrentHeight > highestHeight {
				highestHeight = executor.BTCClients[idx].CurrentHeight
				highestIdx = idx
			}
		}

		//if executor.BTCClients[executor.clientIdx].CurrentHeight+FallBehindThreshold < highestHeight {
		if highestHeight > executor.HighestHeight {
			common.Logger.Infof("new height:" + Int64ToString(highestHeight))

			executor.mutex.Lock()
			executor.clientIdx = highestIdx
			executor.HighestHeight = highestHeight
			executor.mutex.Unlock()
		}
		time.Sleep(time.Duration(executor.Config.BTCConfig.SleepSecond) * time.Second)
	}
}
