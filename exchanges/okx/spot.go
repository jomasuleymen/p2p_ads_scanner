package okx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jomasuleymen/p2p_mining/exchange"

	"github.com/jomasuleymen/p2p_mining/utils"
)

type spotDataFetcher struct{}

func (spot *spotDataFetcher) GetExchange() *exchange.Exchange {
	return instance
}

func (spot *spotDataFetcher) GetSymbol(base, quote string) string {
	return fmt.Sprintf(`%s-%s`, base, quote)
}

func (spot *spotDataFetcher) BookTickersHttpRequest() (*http.Request, error) {
	var bookTickersUrl = "https://www.okx.com/api/v5/market/tickers?instType=SPOT"
	req, err := http.NewRequest(http.MethodGet, bookTickersUrl, nil)

	return req, err
}

type bookTickerT struct {
	Symbol   string `json:"instId"`
	BidPrice string `json:"bidPx"`
	AskPrice string `json:"askPx"`
}

type bookTickersResponse struct {
	Data []*bookTickerT `json:"data"`
}

func (bookTicker *bookTickerT) GetSymbol() string {
	return bookTicker.Symbol
}

func (bookTicker *bookTickerT) GetBidPrice() float64 {
	return utils.ParseFloat(bookTicker.BidPrice)
}

func (bookTicker *bookTickerT) GetAskPrice() float64 {
	return utils.ParseFloat(bookTicker.AskPrice)
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
	Data []*struct {
		Symbol string `json:"instId"`
	} `json:"data"`
}

func (spot *spotDataFetcher) SpotSymbolsHttpRequest() (*http.Request, error) {
	var exchangeInformationUrl = "https://www.okx.com/api/v5/market/tickers?instType=SPOT"
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
