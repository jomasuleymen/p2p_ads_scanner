package okx

import (
	"encoding/json"
	"net/http"

	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/types"
)

var (
	p2pUrl = "https://www.okx.com/v3/c2c/tradingOrders/books"
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
	side := fetcher.GetFormalTradeSide(tradeSide)

	req, err := http.NewRequest(http.MethodGet, p2pUrl, nil)

	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("x-cdn", "https://okx.com")
	req.Header.Set("x-utc", "6")
	req.Header.Set("sec-fetch-site", "same-origin")

	query := req.URL.Query()
	query.Set("side", side)
	query.Set("quoteCurrency", fiat)
	query.Set("baseCurrency", asset)
	query.Set("paymentMethod", "all")
	query.Set("userType", "all")
	query.Set("showTrade", "false")
	query.Set("showFollow", "false")
	query.Set("showAlreadyTraded", "false")
	query.Set("isAbleFilter", "false")

	/* on OKX trade sides are reversed */
	if side == "sell" {
		query.Set("sortType", "price_asc")
	} else {
		query.Set("sortType", "price_desc")
	}

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
	var err = json.Unmarshal(resRaw, &body)

	var trades []*advertise

	if err == nil && body.Data != nil {
		if len(body.Data.Buy) > 0 {
			trades = body.Data.Buy
		} else {
			trades = body.Data.Sell
		}
	}

	advertises := make([]exchange.IP2PAdvertisement, 0, len(trades))

	for _, ads := range trades {
		advertises = append(advertises, ads)
	}

	return &exchange.P2PAdsInfo{
		Advertisements: advertises,
		HasNextPage:    false,
	}
}

func (fetcher *adsFetcher) CheckIfCurrencyPairSupport(fiat string, asset string) bool {
	return isCurrencyPairSupport(fiat, asset)
}
