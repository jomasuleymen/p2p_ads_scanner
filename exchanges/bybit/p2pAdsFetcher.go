package bybit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/types"
)

var (
	p2pUrl  = "https://api2.bybit.com/fiat/otc/item/online"
	adsRows = 150
)

type adsFetcher struct{}

func (fetcher *adsFetcher) GetStartPage() int {
	return 1
}

type p2pRequestBody struct {
	TokenID    string `json:"tokenId"`
	CurrencyID string `json:"currencyId"`
	Side       string `json:"side"`
	Size       string `json:"size"`
	Page       string `json:"page"`
}

func (fetcher *adsFetcher) PrepareHttpRequest(
	fiat string,
	asset string,
	tradeSide types.TradeSide,
	page int,
) (*http.Request, error) {
	reqTradeType := fetcher.GetFormalTradeSide(tradeSide)

	body, _ := json.Marshal(p2pRequestBody{
		TokenID:    asset,
		CurrencyID: fiat,
		Side:       reqTradeType,
		Size:       strconv.Itoa(adsRows),
		Page:       strconv.Itoa(page),
	})

	req, err := http.NewRequest(http.MethodPost, p2pUrl, bytes.NewBuffer(body))

	req.Header.Set("content-type", "application/json")

	return req, err
}

func (fetcher *adsFetcher) GetFormalTradeSide(tradeSide types.TradeSide) string {
	if tradeSide == types.SELL {
		return "0"
	}
	return "1"
}

func (fetcher *adsFetcher) ParseAdvertisements(resRaw []byte) *exchange.P2PAdsInfo {
	var body p2pResponseBody
	var _ = json.Unmarshal(resRaw, &body)

	advertises := make([]exchange.IP2PAdvertisement, 0, len(body.Result.Advertises))

	for _, ads := range body.Result.Advertises {
		advertises = append(advertises, ads)
	}
	return &exchange.P2PAdsInfo{
		Advertisements: advertises,
		HasNextPage:    len(body.Result.Advertises) >= adsRows,
	}
}

func (fetcher *adsFetcher) CheckIfCurrencyPairSupport(fiat string, asset string) bool {
	return isCurrencyPairSupport(fiat, asset)
}
