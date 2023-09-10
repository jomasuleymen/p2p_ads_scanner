package huobi

import (
	"strconv"

	"github.com/jomasuleymen/p2p_mining/utils"
)

type responseBody struct {
	Data      []*advertise `json:"data"`
	CurrPage  int          `json:"currPage"`
	TotalPage int          `json:"totalPage"`
}

type advertise struct {
	ID         int    `json:"id"`
	UID        int    `json:"uid"`
	UserName   string `json:"userName"`
	Currency   int    `json:"currency"`
	PayMethods []*struct {
		Name string `json:"name"`
	} `json:"payMethods"`
	MinTradeLimit     string `json:"minTradeLimit"`
	MaxTradeLimit     string `json:"maxTradeLimit"`
	Price             string `json:"price"`
	TradeCount        string `json:"tradeCount"`
	IsOnline          bool   `json:"isOnline"`
	TradeMonthTimes   int    `json:"tradeMonthTimes"`
	OrderCompleteRate string `json:"orderCompleteRate"`
}

func (ads *advertise) GetAdvertisementId() string {
	return strconv.Itoa(ads.ID)
}

func (ads *advertise) GetPrice() float64 {
	return utils.ParseFloat(ads.Price)
}
func (ads *advertise) GetMinLimit() float64 {
	return utils.ParseFloat(ads.MinTradeLimit)
}

func (ads *advertise) GetMaxLimit() float64 {
	return utils.ParseFloat(ads.MaxTradeLimit)
}

func (ads *advertise) GetAvailableAssetNum() float64 {
	return utils.ParseFloat(ads.TradeCount)
}
func (ads *advertise) GetPaymentMethods() []string {
	var filtered []string

	for _, paymentMethod := range ads.PayMethods {
		filtered = append(filtered, paymentMethod.Name)
	}

	return filtered
}

func (ads *advertise) GetAdvertiserId() string {
	return strconv.Itoa(ads.UID)
}
func (ads *advertise) GetUserName() string {
	return ads.UserName
}

func (ads *advertise) GetUserMonthOrderCount() int {
	return ads.TradeMonthTimes
}

func (ads *advertise) GetUserMonthClosedPercent() int {
	return int(utils.ParseFloat(ads.OrderCompleteRate))
}
