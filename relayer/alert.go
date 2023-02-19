package relayer

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/coredao-org/btc-relayer/common"
	util "github.com/coredao-org/btc-relayer/config"
)

const (
	RetryInterval = 5 * time.Second
)

func (r *Relayer) alert() {
	if !r.cfg.AlertConfig.EnableAlert {
		return
	}
	balanceThreshold, err := decimal.NewFromString(r.cfg.AlertConfig.BalanceThreshold)
	if err != nil {
		panic(err)
	}
	for {
		balance, err := r.coreExecutor.GetRelayerBalance()
		if err != nil {
			time.Sleep(RetryInterval)
			continue
		} else {
			balance, err := decimal.NewFromString(balance.String())
			if err != nil {
				common.Logger.Error(err.Error())
			}
			if r.cfg.AlertConfig.EnableHeartBeat {
				util.SendTelegramMessage(r.cfg.AlertConfig.Identity, r.cfg.AlertConfig.TelegramBotId, r.cfg.AlertConfig.TelegramChatId, fmt.Sprintf("Info: heartbeat message: relayer balance: %s", balance.String()))
			}
			if balance.Cmp(balanceThreshold) <= 0 {
				msg := fmt.Sprintf("Alert: btc-relayer balance (%s:Core) on Core Chain is less than threshold (%s:Core)",
					balance.Div(decimal.NewFromInt(1e18)).String(), balanceThreshold.Div(decimal.NewFromInt(1e18)).String())
				util.SendTelegramMessage(r.cfg.AlertConfig.Identity, r.cfg.AlertConfig.TelegramBotId, r.cfg.AlertConfig.TelegramChatId, msg)
			}
		}

		time.Sleep(time.Duration(r.cfg.AlertConfig.Interval) * time.Second)
	}
}
