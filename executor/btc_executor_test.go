package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBTCExecutor_GetLatestBlockHeight(t *testing.T) {
	executor, err := NewBTCExecutor(cfg)
	require.NoError(t, err)

	height, err := executor.GetLatestBlockHeight(executor.GetClient())
	require.NotNilf(t, height, "error")
}

func TestGetBlockHash(t *testing.T) {
	executor, err := NewBTCExecutor(cfg)
	require.NoError(t, err)

	hash, err := executor.GetBlockHash(executor.GetClient(), 0)
	require.NotNilf(t, hash, "error")
}

func TestGetBlock(t *testing.T) {

	BTCExecutor, err := NewBTCExecutor(cfg)
	require.NoError(t, err)

	height, err := BTCExecutor.GetLatestBlockHeight(BTCExecutor.GetClient())
	require.NoError(t, err)
	require.Greaterf(t, height, int64(0), "must be greater then 0")

	hash, err := BTCExecutor.GetBlockHash(BTCExecutor.GetClient(), height)
	require.NoError(t, err)

	block, err := BTCExecutor.GetBlock(BTCExecutor.GetClient(), hash)
	require.NotNilf(t, block, "error")
}
