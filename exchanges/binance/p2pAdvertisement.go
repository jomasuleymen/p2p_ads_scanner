package binance

import (
	"github.com/jomasuleymen/p2p_mining/utils"
)

type responseBody struct {
	Data []*advertise `json:"data"`
}

type advertise struct {
	Ads *struct {
		AdvNo                string `json:"advNo"`
		TradeType            string `json:"tradeType"`
		Asset                string `json:"asset"`
		FiatUnit             string `json:"fiatUnit"`
		Price                string `json:"price"`
		SurplusAmount        string `json:"surplusAmount"`
		MaxSingleTransAmount string `json:"maxSingleTransAmount"`
		MinSingleTransAmount string `json:"minSingleTransAmount"`
		TradeMethods         []*struct {
			TradeMethodName string `json:"tradeMethodName"`
		} `json:"tradeMethods"`
		DynamicMaxSingleTransAmount   string `json:"dynamicMaxSingleTransAmount"`
		MinSingleTransQuantity        string `json:"minSingleTransQuantity"`
		MaxSingleTransQuantity        string `json:"maxSingleTransQuantity"`
		DynamicMaxSingleTransQuantity string `json:"dynamicMaxSingleTransQuantity"`
		TradableQuantity              string `json:"tradableQuantity"`
	} `json:"adv"`
	Advertiser *advertiser `json:"advertiser"`
}

type advertiser struct {
	UserNo          string  `json:"userNo"`
	NickName        string  `json:"nickName"`
	MonthOrderCount float64 `json:"monthOrderCount"`
	MonthFinishRate float64 `json:"monthFinishRate"`
	UserType        string  `json:"userType"`
	UserGrade       int     `json:"userGrade"`
}

func (ads *advertise) GetAdvertisementId() string {
	return ads.Ads.AdvNo
}

func (ads *advertise) GetPrice() float64 {
	return utils.ParseFloat(ads.Ads.Price)
}

func (ads *advertise) GetMinLimit() float64 {
	return utils.ParseFloat(ads.Ads.MinSingleTransAmount)
}

func (ads *advertise) GetMaxLimit() float64 {
	return utils.ParseFloat(ads.Ads.DynamicMaxSingleTransAmount)
}

func (ads *advertise) GetAvailableAssetNum() float64 {
	return utils.ParseFloat(ads.Ads.TradableQuantity)
}
func (ads *advertise) GetPaymentMethods() []string {
	var paymentMethods []string

	for _, paymentMethod := range ads.Ads.TradeMethods {
		if paymentMethod.TradeMethodName != "" {
			paymentMethods = append(paymentMethods, paymentMethod.TradeMethodName)
		}
	}

	return paymentMethods
}

func (ads *advertise) GetAdvertiserId() string {
	return ads.Advertiser.UserNo
}
func (ads *advertise) GetUserName() string {
	return ads.Advertiser.NickName
}

func (ads *advertise) GetUserMonthOrderCount() int {
	return int(ads.Advertiser.MonthOrderCount)
}

func (ads *advertise) GetUserMonthClosedPercent() int {
	return int(ads.Advertiser.MonthFinishRate * 100)
}
