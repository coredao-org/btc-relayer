package common

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

type Task struct {
	Height    int64
	BlockHash *chainhash.Hash
	BLOCK     *wire.MsgBlock
}

type TaskSet struct {
	TaskList []Task
}
