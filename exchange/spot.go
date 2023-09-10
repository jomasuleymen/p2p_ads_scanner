package exchange

import (
	"net/http"
)

type ISpot interface {
	GetExchange() *Exchange
	GetSymbol(asset1, asset2 string) string
	IBookTickerFetcher
	ISpotSymbolsFetcher
}

type IBookTickerFetcher interface {
	GetExchange() *Exchange
	BookTickersHttpRequest() (*http.Request, error)
	GetBookTickersFromJsonRaw(resRaw []byte) []IBookTicker
}

type ISpotSymbolsFetcher interface {
	GetExchange() *Exchange
	SpotSymbolsHttpRequest() (*http.Request, error)
	GetSpotSymbolsFromJsonRaw(resRaw []byte) []string
}

type IBookTicker interface {
	GetSymbol() string
	GetBidPrice() float64
	GetAskPrice() float64
}
