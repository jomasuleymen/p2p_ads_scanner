package huobi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jomasuleymen/p2p_mining/exchange"
)

type spotDataFetcher struct{}

func (spot *spotDataFetcher) GetExchange() *exchange.Exchange {
	return instance
}

func (spot *spotDataFetcher) GetSymbol(base, quote string) string {
	return strings.ToLower(fmt.Sprintf(`%s%s`, base, quote))
}

func (spot *spotDataFetcher) BookTickersHttpRequest() (*http.Request, error) {
	var bookTickersUrl = "https://api.huobi.pro/market/tickers"
	req, err := http.NewRequest(http.MethodGet, bookTickersUrl, nil)

	return req, err
}

type bookTickerT struct {
	Symbol   string  `json:"symbol"`
	BidPrice float64 `json:"bid"`
	AskPrice float64 `json:"ask"`
}

type bookTickersResponse struct {
	Data []*bookTickerT `json:"data"`
}

func (bookTicker *bookTickerT) GetSymbol() string {
	return bookTicker.Symbol
}

func (bookTicker *bookTickerT) GetBidPrice() float64 {
	return bookTicker.BidPrice
}

func (bookTicker *bookTickerT) GetAskPrice() float64 {
	return bookTicker.AskPrice
}

func (spot *spotDataFetcher) GetBookTickersFromJsonRaw(resRaw []byte) []exchange.IBookTicker {
	var data *bookTickersResponse

	var _ = json.Unmarshal(resRaw, &data)

	var bookTickers []exchange.IBookTicker

	for _, bookTicker := range data.Data {
		bookTickers = append(bookTickers, bookTicker)
	}

	return bookTickers
}

type spotSymbolT struct {
	Data []struct {
		Symbol string `json:"symbol"`
	} `json:"data"`
}

func (spot *spotDataFetcher) SpotSymbolsHttpRequest() (*http.Request, error) {
	var exchangeInformationUrl = "https://api.huobi.pro/market/tickers"
	req, err := http.NewRequest(http.MethodGet, exchangeInformationUrl, nil)

	return req, err
}

func (spot *spotDataFetcher) GetSpotSymbolsFromJsonRaw(resRaw []byte) []string {
	var data *spotSymbolT

	var _ = json.Unmarshal(resRaw, &data)

	var symbols []string

	for _, symbol := range data.Data {
		symbols = append(symbols, symbol.Symbol)
	}

	return symbols
}
