package relayer

import (
	config "github.com/coredao-org/btc-relayer/config"
	"github.com/coredao-org/btc-relayer/executor"
)

type Relayer struct {
	cfg          *config.Config
	btcExecutor  *executor.BTCExecutor
	coreExecutor *executor.COREExecutor
}

func NewRelayer(cfg *config.Config, BTCExecutor *executor.BTCExecutor, coreExecutor *executor.COREExecutor) *Relayer {
	return &Relayer{
		cfg:          cfg,
		btcExecutor:  BTCExecutor,
		coreExecutor: coreExecutor,
	}
}

func (r *Relayer) Start() {

	//register relayer
	r.registerRelayerHub()

	go r.RelayerCompetitionDaemon()

	go r.btcExecutor.UpdateClients()
	go r.coreExecutor.UpdateClients()

	go r.alert()
}
