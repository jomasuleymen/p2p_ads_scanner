package pexpay

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/types"
)

var (
	p2pUrl  = "https://www.pexpay.com/bapi/c2c/v1/friendly/c2c/ad/search"
	adsRows = 20
)

type adsFetcher struct{}

func (fetcher *adsFetcher) GetStartPage() int {
	return 1
}

type p2pRequestPayload struct {
	Page          int    `json:"page"`
	Rows          int    `json:"rows"`
	Asset         string `json:"asset"`
	Fiat          string `json:"fiat"`
	MerchantCheck bool   `json:"merchantCheck"`
	TradeType     string `json:"tradeType"`
}

func (fetcher *adsFetcher) PrepareHttpRequest(
	fiat string,
	asset string,
	tradeSide types.TradeSide,
	page int,
) (*http.Request, error) {
	reqTradeType := fetcher.GetFormalTradeSide(tradeSide)

	payload := &p2pRequestPayload{
		Page:          page,
		Rows:          adsRows,
		Asset:         asset,
		Fiat:          fiat,
		MerchantCheck: false,
		TradeType:     reqTradeType,
	}

	jsonStr, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, p2pUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("clienttype", "web")
	req.Header.Set("Content-Type", "application/json")

	return req, err
}

func (fetcher *adsFetcher) GetFormalTradeSide(tradeSide types.TradeSide) string {
	return string(tradeSide)
}

func (fetcher *adsFetcher) ParseAdvertisements(resRaw []byte) *exchange.P2PAdsInfo {
	var body resBody
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
