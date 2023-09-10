package pexpay

import (
	"github.com/jomasuleymen/p2p_mining/exchange"
)

var exchangeName string = "Pexpay"

var instance *exchange.Exchange

func GetInstance() *exchange.Exchange {
	if instance == nil {
		instance = &exchange.Exchange{
			Name:          exchangeName,
			Spot:          nil,
			P2PAdsFetcher: &adsFetcher{},
		}
	}

	return instance
}
