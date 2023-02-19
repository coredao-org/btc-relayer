package relayer

import (
	"time"

	"github.com/coredao-org/btc-relayer/common"
)

func (r *Relayer) registerRelayerHub() {
	isRelayer, err := r.coreExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if isRelayer {
		common.Logger.Info("This relayer has already been registered")
		return
	}

	common.Logger.Info("Register this relayer to RelayerHub")
	_, err = r.coreExecutor.RegisterRelayer()
	if err != nil {
		panic(err)
	}
	common.Logger.Info("Waiting for registration tx to be confirmed")
	time.Sleep(20 * time.Second)

	isRelayer, err = r.coreExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if !isRelayer {
		panic("Failed to register relayer")
	}
}
