package exchange

import (
	"net/http"

	"github.com/jomasuleymen/p2p_mining/types"
)

const (
	Maker = "maker"
	Taker = "taker"
)

type IP2PAdsFetcher interface {
	GetStartPage() int
	GetFormalTradeSide(tradeSide types.TradeSide) string
	PrepareHttpRequest(
		fiat string,
		asset string,
		tradeSide types.TradeSide,
		page int,
	) (*http.Request, error)
	ParseAdvertisements(resRaw []byte) *P2PAdsInfo
	CheckIfCurrencyPairSupport(fiat string, asset string) bool
}

type P2PAdsInfo struct {
	Advertisements []IP2PAdvertisement
	HasNextPage    bool
}

type IP2PAdvertisement interface {
	IP2PAdvertisementData
	IP2PAdvertiser
}

type IP2PAdvertisementData interface {
	GetAdvertisementId() string
	GetPrice() float64
	GetMinLimit() float64
	GetMaxLimit() float64
	GetAvailableAssetNum() float64
	GetPaymentMethods() []string
}

type IP2PAdvertiser interface {
	GetAdvertiserId() string
	GetUserName() string
	GetUserMonthOrderCount() int
	GetUserMonthClosedPercent() int
}

type P2PSchemeFilter struct {
	Fiat   int `json:"fiat"`
	Assets struct {
		Selected int   `json:"selected"`
		Declined []int `json:"declined"`
	} `json:"assets"`
	Exchanges struct {
		Declined []int `json:"declined"`
	} `json:"exchanges"`
	Actions struct {
		BuyAs  string `json:"buyAs"`
		SellAs string `json:"sellAs"`
	} `json:"actions"`
	Payments struct {
		WeBuy    []int `json:"weBuy"`
		WeSell   []int `json:"weSell"`
		Declined []int `json:"declined"`
	} `json:"payments"`
	User struct {
		MinOrder  int     `json:"minOrder"`
		MinRate   float64 `json:"minRate"`
		IgnoreIds []int   `json:"ignoreIds"`
	} `json:"user"`
	SpreadPercent struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"spreadPercent"`
}
