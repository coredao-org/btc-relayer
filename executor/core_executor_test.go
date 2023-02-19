package executor

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/coredao-org/btcpowermirror/lightmirror"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"

	relayercommon "github.com/coredao-org/btc-relayer/common"

	config "github.com/coredao-org/btc-relayer/config"
	"github.com/stretchr/testify/require"
)

const (
	provider   = "http://{coreRpcAddress}"
	privateKey = ""
)

var (
	cfg = &config.Config{
		CrossChainConfig: config.CrossChainConfig{
			RecursionHeight: 10,
		},
		BTCConfig: config.BTCConfig{
			RpcAddrs: []config.BTCRpcAddrs{{Host: "btc_address}", User: "", Pass: ""}},
		},
		COREConfig: config.COREConfig{
			GasLimit:   4700000,
			Providers:  []string{provider},
			PrivateKey: privateKey,
		},
	}
)

func TestCOREExecutor_NewBTCExecutor(t *testing.T) {
	BTCExecutor, err := NewBTCExecutor(cfg)
	require.NoError(t, err)
	require.NotNilf(t, BTCExecutor, "error")
}

func TestCOREExecutor_GetLatestBlockHeight(t *testing.T) {
	executor, err := NewCOREExecutor(cfg)
	require.NoError(t, err)

	height, err := executor.GetLatestBlockHeight(executor.GetClient())
	require.NoError(t, err)
	require.Greaterf(t, height, int64(0), "")
}

func TestCOREExecutor_UpdateClients(t *testing.T) {
	executor, err := NewCOREExecutor(cfg)
	require.NoError(t, err)

	executor.UpdateClients()
	require.NoError(t, err)
}

func TestCOREExecutor_RegisterRelayer(t *testing.T) {
	executor, err := NewCOREExecutor(cfg)
	require.NoError(t, err)

	txHash, err := executor.RegisterRelayer()
	require.NoError(t, err)
	require.NotNil(t, txHash)
}

func TestCOREExecutor_IsRelayer(t *testing.T) {
	executor, err := NewCOREExecutor(cfg)
	require.NoError(t, err)

	isRelayer, err := executor.IsRelayer()
	require.NoError(t, err)
	require.Equal(t, isRelayer, true, "")
}

func TestCOREExecutor_CheckBlockRelayed(t *testing.T) {
	executor, err := NewCOREExecutor(cfg)
	require.NoError(t, err)

	result, err := executor.CheckBlockRelayed(&chainhash.Hash{})
	require.NoError(t, err)
	require.Equal(t, false, result, "")
}

func TestCOREExecutor_SyncBTCLightMirror(t *testing.T) {
	BTCExecutor, err := NewBTCExecutor(cfg)
	require.NoError(t, err)
	executor, err := NewCOREExecutor(cfg)
	require.NoError(t, err)

	hash, err := BTCExecutor.GetBlockHash(BTCExecutor.GetClient(), 717696)
	require.NoError(t, err)

	block, err := BTCExecutor.GetBlock(BTCExecutor.GetClient(), hash)
	require.NoError(t, err)

	task := relayercommon.Task{BLOCK: block, BlockHash: hash, Height: 717696}
	txHash, err := executor.SyncBTCLightMirror(&task)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestCOREExecutor_testGeneratorPublicKey(t *testing.T) {
	privateKey := ""
	privKey, err := crypto.HexToECDSA(privateKey)
	require.NoError(t, err)

	publicKey := privKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	pub := hex.EncodeToString(crypto.FromECDSAPub(publicKeyECDSA))
	require.NotNil(t, pub)
}

func TestCOREExecutor_testCore2(t *testing.T) {
	btcExecutor, err := NewBTCExecutor(cfg)
	btcHash, err := btcExecutor.GetClient().GetBlockHash(766080)
	if err != nil {
		return
	}

	log.Println(btcHash.String())

	header, err := btcExecutor.GetClient().GetBlock(btcHash)

	var b bytes.Buffer
	header.Serialize(&b)
	bb := b.Bytes()
	ss := hex.EncodeToString(bb)
	log.Println(ss)

	printMirror(NewBtcLightMirror(header))

	require.NotNil(t, ss)
}

func printMirror(mirror *lightmirror.BtcLightMirrorV2) {
	var b bytes.Buffer
	mirror.Serialize(&b)
	bb := b.Bytes()

	ss := hex.EncodeToString(bb)
	log.Println(ss)
}
