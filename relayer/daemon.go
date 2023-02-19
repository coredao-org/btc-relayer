package relayer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/coredao-org/btc-relayer/common"
	"github.com/coredao-org/btc-relayer/executor"
)

var (
	ZERO_HASH = chainhash.Hash{00000000000000000000000000000000}
)

func (r *Relayer) getLatestHeight() uint64 {
	height, err := r.btcExecutor.GetLatestBlockHeight(r.btcExecutor.GetClient())
	if err != nil {
		common.Logger.Errorf("Query latest height error: %s", err.Error())
		return 0
	}
	return uint64(height)
}

func (r *Relayer) getLastRelayHeight() (int64, error) {
	//last relayed btc block hash
	chainTip, err := r.coreExecutor.GetChainTip()
	if err != nil {
		return 0, err
	}

	chainTip = executor.RevertHash(chainTip)

	blockHeaderVerbose, err := r.btcExecutor.GetClient().GetBlockHeaderVerbose(chainTip)
	if err != nil {
		return 0, err
	}

	height := int64(blockHeaderVerbose.Height)
	blockHash, err := r.btcExecutor.GetClient().GetBlockHash(height)
	if err != nil {
		return 0, err
	}
	blockHeaderVerboseNew, err := r.btcExecutor.GetClient().GetBlockHeaderVerbose(blockHash)

	//Forked, need to push backwards
	if chainTip.String() != blockHeaderVerboseNew.Hash {
		lastHeight, err := r.recursionGetLastHeight(height)
		if err == nil {
			height = lastHeight
		}
	}

	return height, nil
}

func (r *Relayer) recursionGetLastHeight(height int64) (int64, error) {
	for {
		blockHash, err := r.btcExecutor.GetClient().GetBlockHash(height)
		if err != nil {
			return height, err
		}

		relayed, err := r.coreExecutor.CheckBlockRelayed(blockHash)
		if err != nil {
			return height, err
		}

		if relayed {
			return height, nil
		}

		height -= 1
	}
}

func (r *Relayer) recursionUnCommitTasks(height int64, lastBlockHash *chainhash.Hash) (*common.TaskSet, error) {
	var taskSet common.TaskSet

	//get HighestHeight block hash
	blockHash, err := r.btcExecutor.GetClient().GetBlockHash(r.btcExecutor.HighestHeight)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	for {
		if blockHash.IsEqual(lastBlockHash) {
			break
		}

		//check if this block is relayed
		relayed, err := r.CheckBlockRelayed(blockHash)
		if err != nil {
			return nil, fmt.Errorf("error")
		}

		if relayed {
			break
		}

		//get block
		block, err := r.btcExecutor.GetClient().GetBlock(blockHash)

		json, _ := json.Marshal(block)
		print(json)

		if err != nil {
			break
		}

		//new block
		task := common.Task{
			Height:    height,
			BlockHash: blockHash,
			BLOCK:     block,
		}

		//append to head
		taskSet.TaskList = append([]common.Task{task}, taskSet.TaskList...)

		if r.cfg.CrossChainConfig.RecursionHeight > 0 && int64(len(taskSet.TaskList)) >= r.cfg.CrossChainConfig.RecursionHeight {
			break
		}

		//get pre block hash
		blockHash = &block.Header.PrevBlock
		//Genesis Block
		if blockHash.IsEqual(&ZERO_HASH) {
			break
		}

		height--
	}

	return &taskSet, nil
}

func (r *Relayer) RelayerCompetitionDaemon() {

	//var err error
	common.Logger.Info("Start relayer daemon")

	for {
		//no new block, sleep
		if r.btcExecutor.HighestHeight == (int64(0)) {
			time.Sleep(time.Second)
			continue
		}

		lastRelayHeight, err := r.getLastRelayHeight()

		if err != nil {
			continue
		}

		//no new block, sleep 1s
		if lastRelayHeight == r.btcExecutor.HighestHeight {
			common.Logger.Infof("no new block, current height:" + executor.Int64ToString(lastRelayHeight))
			time.Sleep(1 * time.Second)
			continue
		}

		common.Logger.Infof("find last relayed height:" + executor.Int64ToString(lastRelayHeight))

		for i := lastRelayHeight + 1; i <= r.btcExecutor.HighestHeight; {
			common.Logger.Infof("start relaying, height:" + executor.Int64ToString(i))

			_, err := r.DoRelayWithHeight(i)
			if err == nil {
				common.Logger.Infof("successfully relayed, height:" + executor.Int64ToString(i))
				i++
				continue
			} else {
				time.Sleep(3 * time.Second)
				common.Logger.Infof("relay failed, height:"+executor.Int64ToString(i), err)
			}
		}
	}
}

/**
do relay
*/
func (r *Relayer) DoRelayWithHeight(blockHeight int64) (bool, error) {
	blockHash, err := r.btcExecutor.GetClient().GetBlockHash(blockHeight)
	if err != nil {
		return false, err
	}

	//check if this block is relayed
	relayed, err := r.CheckBlockRelayed(blockHash)
	if err != nil {
		return false, err
	}

	if relayed {
		common.Logger.Infof("block is relayed, height:" + executor.Int64ToString(blockHeight))
		return true, nil
	}

	//get block
	block, err := r.btcExecutor.GetClient().GetBlock(blockHash)

	if err != nil {
		return false, err
	}

	//new block
	task := common.Task{
		Height:    blockHeight,
		BlockHash: blockHash,
		BLOCK:     block,
	}

	err = r.doRelay(&task)

	return err == nil, err
}
func (r *Relayer) doRelay(task *common.Task) error {
	_, err := r.coreExecutor.SyncBTCLightMirror(task)

	return err
}

func (r *Relayer) CheckBlockRelayed(blockHash *chainhash.Hash) (bool, error) {
	//check if this block if relayed
	checkResult, err := r.coreExecutor.CheckBlockRelayed(blockHash)
	if err != nil {
		return true, fmt.Errorf("error")
	}

	return checkResult, nil
}
