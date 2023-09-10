package huobi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/types"
)

var (
	p2pUrl = "https://otc-akm.huobi.com/v1/data/trade-market"
	doc    = map[string]map[string]string{
		"fiats": {
			"USD": "2",
			"RUB": "11",
			"KZT": "57",
		},
		"assets": {
			"BTC":  "1",
			"USDT": "2",
			"ETH":  "3",
		},
	}
)

type adsFetcher struct{}

func (fetcher *adsFetcher) GetStartPage() int {
	return 1
}

func (fetcher *adsFetcher) PrepareHttpRequest(
	fiat string,
	asset string,
	tradeSide types.TradeSide,
	page int,
) (*http.Request, error) {
	reqTradeType := fetcher.GetFormalTradeSide(tradeSide)

	req, err := http.NewRequest(http.MethodGet, p2pUrl, nil)
	query := req.URL.Query()
	query.Set("coinId", doc["assets"][asset])
	query.Set("currency", doc["fiats"][fiat])
	query.Set("tradeType", reqTradeType)
	query.Set("currPage", strconv.Itoa(page))
	query.Set("payMethod", "0")
	query.Set("acceptOrder", "0")
	query.Set("blockType", "general")
	query.Set("range", "0")
	query.Set("online", "1")
	query.Set("isFollowed", "false")
	query.Set("onlyTradable", "true")
	req.URL.RawQuery = query.Encode()

	return req, err
}

func (fetcher *adsFetcher) GetFormalTradeSide(tradeSide types.TradeSide) string {
	if tradeSide == types.BUY {
		return "sell"
	}

	return "buy"
}

func (fetcher *adsFetcher) ParseAdvertisements(resRaw []byte) *exchange.P2PAdsInfo {
	var body responseBody
	var _ = json.Unmarshal(resRaw, &body)

	advertises := make([]exchange.IP2PAdvertisement, 0, len(body.Data))
	hasNext := true

	for _, ads := range body.Data {
		if !ads.IsOnline {
			hasNext = false
			break
		}

		advertises = append(advertises, ads)
	}

	if hasNext {
		hasNext = body.CurrPage < body.TotalPage
	}

	return &exchange.P2PAdsInfo{
		Advertisements: advertises,
		HasNextPage:    hasNext,
	}
}

func (fetcher *adsFetcher) CheckIfCurrencyPairSupport(fiat string, asset string) bool {
	return isCurrencyPairSupport(fiat, asset)
}
