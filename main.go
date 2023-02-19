package main

import (
	"fmt"
	"github.com/coredao-org/btc-relayer/common"
	config "github.com/coredao-org/btc-relayer/config"
	"github.com/coredao-org/btc-relayer/executor"
	"github.com/coredao-org/btc-relayer/relayer"
)

const (
	flagConfigPath = "config.json"
)

func printUsage() {
	fmt.Print("usage: ./btc-relayer --config-path configFile\n")
}

/**
init cfg
*/
func initCfg() *config.Config {

	var cfg *config.Config

	cfg = config.ParseConfigFromFile(flagConfigPath)

	if cfg == nil {
		common.Logger.Infof("failed to get configuration")
	}

	return cfg
}

func main() {

	//init cfg
	cfg := initCfg()

	//init logger
	common.InitLogger(&cfg.LogConfig)

	//init executors
	btcExecutor, err := executor.NewBTCExecutor(cfg)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	coreExecutor, err := executor.NewCOREExecutor(cfg)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	relayerInstance := relayer.NewRelayer(cfg, btcExecutor, coreExecutor)

	common.Logger.Info("Starting relayer")
	relayerInstance.Start()

	select {}
}
