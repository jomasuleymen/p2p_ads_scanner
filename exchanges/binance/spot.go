package binance

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
	return fmt.Sprintf(`%s%s`, base, quote)
}

func (spot *spotDataFetcher) BookTickersHttpRequest() (*http.Request, error) {
	var bookTickersUrl = "https://api.binance.com/api/v3/ticker/bookTicker"
	req, err := http.NewRequest(http.MethodGet, bookTickersUrl, nil)

	return req, err
}

type bookTickerT struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	AskPrice string `json:"askPrice"`
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
	var allBookTickers []*bookTickerT

	var _ = json.Unmarshal(resRaw, &allBookTickers)

	var bookTickers []exchange.IBookTicker

	for _, bookTicker := range allBookTickers {
		bookTickers = append(bookTickers, bookTicker)
	}

	return bookTickers
}

type spotSymbolT struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	AskPrice string `json:"askPrice"`
}

func (spot *spotDataFetcher) SpotSymbolsHttpRequest() (*http.Request, error) {
	var exchangeInformationUrl = "https://api.binance.com/api/v3/ticker/bookTicker"
	req, err := http.NewRequest(http.MethodGet, exchangeInformationUrl, nil)

	return req, err
}

func (spot *spotDataFetcher) GetSpotSymbolsFromJsonRaw(resRaw []byte) []string {
	var allSpotSymbols []*spotSymbolT

	var _ = json.Unmarshal(resRaw, &allSpotSymbols)

	var symbols []string

	for _, ticker := range allSpotSymbols {
		if utils.ParseFloat(ticker.BidPrice) != 0 || utils.ParseFloat(ticker.AskPrice) != 0 {
			symbols = append(symbols, ticker.Symbol)
		}
	}

	return symbols
}
