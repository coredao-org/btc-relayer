package executor

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	DefaultGasPrice = 20000000000 // 20 GWei

	FallBehindThreshold          = 5
	DataSeedDenyServiceThreshold = 60
)

var (
	prefixForCrossChainPackageKey = []byte{0x00}
	prefixForSequenceKey          = []byte{0xf0}

	relayerIncentivizeContractAddr = common.HexToAddress("0x0000000000000000000000000000000000001005")
	relayerHubContractAddr         = common.HexToAddress("0x0000000000000000000000000000000000001004")
	crossChainContractAddr         = common.HexToAddress("0x0000000000000000000000000000000000002000")
	pcsAddr                        = common.HexToAddress("0x0000000000000000000000000000000000001003")
)
