package binance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jomasuleymen/p2p_mining/exchange"

	"github.com/jomasuleymen/p2p_mining/types"
)

var (
	p2pUrl  = "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"
	adsRows = 20
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

	body := []byte(fmt.Sprintf(
		`{"fiat":"%s","asset":"%s","tradeType":"%s","page":%d,"rows":%d,"proMerchantAds":false,"payTypes":[],"countries":[],"publisherType":null,"merchantCheck":false}`,
		fiat,
		asset,
		reqTradeType,
		page,
		adsRows,
	))

	req, err := http.NewRequest(http.MethodPost, p2pUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	return req, err
}

func (fetcher *adsFetcher) GetFormalTradeSide(tradeSide types.TradeSide) string {
	return string(tradeSide)
}

func (fetcher *adsFetcher) ParseAdvertisements(resRaw []byte) *exchange.P2PAdsInfo {
	var body responseBody
	var _ = json.Unmarshal(resRaw, &body)

	advertises := make([]exchange.IP2PAdvertisement, 0, len(body.Data))

	for _, ads := range body.Data {
		advertises = append(advertises, ads)
	}

	return &exchange.P2PAdsInfo{
		Advertisements: advertises,
		HasNextPage:    len(body.Data) >= adsRows,
	}
}

func (fetcher *adsFetcher) CheckIfCurrencyPairSupport(fiat string, asset string) bool {
	return isCurrencyPairSupport(fiat, asset)
}
