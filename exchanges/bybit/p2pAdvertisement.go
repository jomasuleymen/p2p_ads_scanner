package bybit

import (
	"github.com/jomasuleymen/p2p_mining/utils"
)

type p2pResponseBody struct {
	Result struct {
		Count      int          `json:"count"`
		Advertises []*advertise `json:"items"`
	} `json:"result"`
}

type advertise struct {
	ID                string   `json:"id"`
	AccountID         string   `json:"accountId"`
	UserID            string   `json:"userId"`
	NickName          string   `json:"nickName"`
	Price             string   `json:"price"`
	LastQuantity      string   `json:"lastQuantity"`
	Quantity          string   `json:"quantity"`
	MinAmount         string   `json:"minAmount"`
	MaxAmount         string   `json:"maxAmount"`
	Payments          []string `json:"payments"`
	RecentOrderNum    int      `json:"recentOrderNum"`
	RecentExecuteRate int      `json:"recentExecuteRate"`
	IsOnline          bool     `json:"isOnline"`
}

func (ads *advertise) GetAdvertisementId() string {
	return ads.ID
}

func (ads *advertise) GetPrice() float64 {
	return utils.ParseFloat(ads.Price)
}
func (ads *advertise) GetMinLimit() float64 {
	return utils.ParseFloat(ads.MinAmount)
}

func (ads *advertise) GetMaxLimit() float64 {
	return utils.ParseFloat(ads.MaxAmount)
}

func (ads *advertise) GetAvailableAssetNum() float64 {
	return utils.ParseFloat(ads.LastQuantity)
}
func (ads *advertise) GetPaymentMethods() []string {
	var filtered []string

	for _, paymentMethodId := range ads.Payments {
		if val, ok := bybitPaymentMethodsMap[paymentMethodId]; ok {
			filtered = append(filtered, val)
		}
	}

	return filtered
}

func (ads *advertise) GetAdvertiserId() string {
	return ads.UserID
}
func (ads *advertise) GetUserName() string {
	return ads.NickName
}

func (ads *advertise) GetUserMonthOrderCount() int {
	return ads.RecentOrderNum
}

func (ads *advertise) GetUserMonthClosedPercent() int {
	return ads.RecentExecuteRate
}
