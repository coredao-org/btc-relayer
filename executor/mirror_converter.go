package executor

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/coredao-org/btcpowermirror/lightmirror"
	"github.com/jinzhu/copier"
)

func NewBtcLightMirror(block *wire.MsgBlock) *lightmirror.BtcLightMirrorV2 {
	return lightmirror.CreateBtcLightMirrorV2(&block.Header,
		block.Transactions[0],
		fillTxHashes(block.Transactions),
	)
}

func fillTxHashes(transactions []*wire.MsgTx) []chainhash.Hash {
	txHashes := make([]chainhash.Hash, len(transactions))

	for i := range transactions {
		copier.Copy(&txHashes[i], transactions[i].TxHash())
	}

	return txHashes
}
