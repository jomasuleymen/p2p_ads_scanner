package bybit

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
	var bookTickersUrl = "https://api.bybit.com/v5/market/tickers?category=spot"
	req, err := http.NewRequest(http.MethodGet, bookTickersUrl, nil)

	return req, err
}

type bookTickerT struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bid1Price"`
	AskPrice string `json:"ask1Price"`
}

type bookTickersResponse struct {
	Result struct {
		List []*bookTickerT `json:"list"`
	} `json:"result"`
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

	for _, bookTicker := range data.Result.List {
		bookTickers = append(bookTickers, bookTicker)
	}

	return bookTickers
}

type spotSymbolT struct {
	Result struct {
		List []struct {
			Symbol string `json:"symbol"`
		} `json:"list"`
	} `json:"result"`
}

func (spot *spotDataFetcher) SpotSymbolsHttpRequest() (*http.Request, error) {
	var exchangeInformationUrl = "https://api.bybit.com/v5/market/tickers?category=spot"
	req, err := http.NewRequest(http.MethodGet, exchangeInformationUrl, nil)

	return req, err
}

func (spot *spotDataFetcher) GetSpotSymbolsFromJsonRaw(resRaw []byte) []string {
	var data *spotSymbolT

	var _ = json.Unmarshal(resRaw, &data)

	var symbols []string

	for _, symbol := range data.Result.List {
		symbols = append(symbols, symbol.Symbol)
	}

	return symbols
}
