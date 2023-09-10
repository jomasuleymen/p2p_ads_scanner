package okx

import "github.com/jomasuleymen/p2p_mining/utils"

type responseBody struct {
	Data *struct {
		Sell []*advertise `json:"sell"`
		Buy  []*advertise `json:"buy"`
	} `json:"data"`
}

type advertise struct {
	ID                     string   `json:"id"`
	NickName               string   `json:"nickName"`
	PublicUserID           string   `json:"publicUserId"`
	CompletedOrderQuantity int      `json:"completedOrderQuantity"`
	CompletedRate          string   `json:"completedRate"`
	PaymentMethods         []string `json:"paymentMethods"`
	Price                  string   `json:"price"`
	QuoteCurrency          string   `json:"quoteCurrency"`
	AvailableAmount        string   `json:"availableAmount"`
	QuoteMaxAmountPerOrder string   `json:"quoteMaxAmountPerOrder"`
	QuoteMinAmountPerOrder string   `json:"quoteMinAmountPerOrder"`
}

func (ads *advertise) GetAdvertisementId() string {
	return ads.ID
}

func (ads *advertise) GetPrice() float64 {
	return utils.ParseFloat(ads.Price)
}
func (ads *advertise) GetMinLimit() float64 {
	return utils.ParseFloat(ads.QuoteMinAmountPerOrder)
}

func (ads *advertise) GetMaxLimit() float64 {
	return utils.ParseFloat(ads.QuoteMaxAmountPerOrder)
}

func (ads *advertise) GetAvailableAssetNum() float64 {
	return utils.ParseFloat(ads.AvailableAmount)
}
func (ads *advertise) GetPaymentMethods() []string {
	return ads.PaymentMethods
}

func (ads *advertise) GetAdvertiserId() string {
	return ads.PublicUserID
}
func (ads *advertise) GetUserName() string {
	return ads.NickName
}

func (ads *advertise) GetUserMonthOrderCount() int {
	return ads.CompletedOrderQuantity
}

func (ads *advertise) GetUserMonthClosedPercent() int {
	return int(utils.ParseFloat(ads.CompletedRate) * 100)
}
