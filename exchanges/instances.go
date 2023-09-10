package exchanges

import (
	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/exchanges/binance"
	"github.com/jomasuleymen/p2p_mining/exchanges/bybit"
	"github.com/jomasuleymen/p2p_mining/exchanges/huobi"
	"github.com/jomasuleymen/p2p_mining/exchanges/okx"
	"github.com/jomasuleymen/p2p_mining/exchanges/pexpay"
)

var AllExchanges = []*exchange.Exchange{
	binance.GetInstance(),
	okx.GetInstance(),
	bybit.GetInstance(),
	huobi.GetInstance(),
	pexpay.GetInstance(),
}
