package examples

import (
	"cryptobot/strategies"
	"github.com/saniales/golang-crypto-trading-bot/environment"
	"github.com/saniales/golang-crypto-trading-bot/exchanges"
	"github.com/sirupsen/logrus"
)

// Websocket strategy
var Websocket = strategies.WebsocketStrategy{
	Model: strategies.StrategyModel{
		Name: "Websocket",
		Setup: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			for _, wrapper := range wrappers {
				err := wrapper.FeedConnect(markets)
				if err == exchanges.ErrWebsocketNotSupported || err == nil {
					continue
				}
				return err
			}
			return nil
		},
		OnUpdate: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			// do something
			return nil
		},
		TearDown: func(wrappers []exchanges.ExchangeWrapper, markets []*environment.Market) error {
			return nil
		},
		OnError: func(err error) {
			logrus.Error(err)
		},
	},
}
