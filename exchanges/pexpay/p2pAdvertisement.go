package pexpay

import "github.com/jomasuleymen/p2p_mining/utils"

type resBody struct {
	Data []*advertise `json:"data"`
}

type advertise struct {
	Adv *struct {
		AdvNo                string `json:"adNo"`
		TradeSide            string `json:"tradeType"`
		Asset                string `json:"asset"`
		FiatUnit             string `json:"fiatCurrency"`
		Price                string `json:"price"`
		MaxSingleTransAmount string `json:"maxSingleTransAmount"`
		MinSingleTransAmount string `json:"minSingleTransAmount"`
		TradeMethods         []*struct {
			TradeMethodName string `json:"tradeMethodName"`
		} `json:"tradeMethods"`
		IsTradable                    bool   `json:"isTradable"`
		DynamicMaxSingleTransAmount   string `json:"dynamicMaxSingleTransAmount"`
		MinSingleTransQuantity        string `json:"minSingleTransQuantity"`
		MaxSingleTransQuantity        string `json:"maxSingleTransQuantity"`
		DynamicMaxSingleTransQuantity string `json:"dynamicMaxSingleTransQuantity"`
		TradableQuantity              string `json:"tradableQuantity"`
	} `json:"adDetailResp"`
	Advertiser *advertiser `json:"advertiserVo"`
}

type advertiser struct {
	UserNo      string `json:"userNo"`
	NickName    string `json:"nickName"`
	StatsDetail struct {
		MonthOrderCount float64 `json:"completedOrderNum"`
		MonthFinishRate float64 `json:"finishRateLatest30day"`
	} `json:"userStatsRet"`
}

func (ads *advertise) GetAdvertisementId() string {
	return ads.Adv.AdvNo
}

func (ads *advertise) GetPrice() float64 {
	return utils.ParseFloat(ads.Adv.Price)
}
func (ads *advertise) GetMinLimit() float64 {
	return utils.ParseFloat(ads.Adv.MinSingleTransAmount)
}

func (ads *advertise) GetMaxLimit() float64 {
	return utils.ParseFloat(ads.Adv.DynamicMaxSingleTransAmount)
}

func (ads *advertise) GetAvailableAssetNum() float64 {
	return utils.ParseFloat(ads.Adv.TradableQuantity)
}
func (ads *advertise) GetPaymentMethods() []string {
	var filtered []string

	for _, paymentMethod := range ads.Adv.TradeMethods {
		if paymentMethod.TradeMethodName != "" {
			filtered = append(filtered, paymentMethod.TradeMethodName)
		}
	}

	return filtered
}

func (ads *advertise) GetAdvertiserId() string {
	return ads.Advertiser.UserNo
}
func (ads *advertise) GetUserName() string {
	return ads.Advertiser.NickName
}

func (ads *advertise) GetUserMonthOrderCount() int {
	return int(ads.Advertiser.StatsDetail.MonthOrderCount)
}

func (ads *advertise) GetUserMonthClosedPercent() int {
	return int(ads.Advertiser.StatsDetail.MonthFinishRate * 100)
}
