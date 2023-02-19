package executor

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"strconv"
)


func RevertHash(blockHash *chainhash.Hash) (*chainhash.Hash) {

	bHash := *blockHash

	for i := 0; i < 16; i++ {
		bHash[i], bHash[32-1-i] = bHash[32-1-i], bHash[i]
	}

	return &bHash
}

func Int64ToString(value int64) string{
	return strconv.FormatInt(value,10)
}
