package models

type AdsPaymentMethods struct {
	AdvertisementID int `gorm:"uniqueIndex:adsPaymentIndex;type:int8;not null;"`
	PaymentMethodID int `gorm:"uniqueIndex:adsPaymentIndex;type:int8;not null;"`
}
