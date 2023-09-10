package models

import (
	"github.com/jomasuleymen/p2p_mining/types"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Advertisement struct {
	ID                 int              `gorm:"primaryKey;type:bigint;autoincrement:true"`
	AdvertisementID    string           `gorm:"index;uniqueIndex:exchange_advertisement_id;type:varchar(255);not null;"`
	ExchangeID         int              `gorm:"index;uniqueIndex:exchange_advertisement_id;type:smallint;not null;"`
	FiatID             int              `gorm:"index;type:smallint;not null;"`
	AssetID            int              `gorm:"index;type:smallint;not null;"`
	AdvertiserID       int              `gorm:"index;type:bigint;not null;"`
	RowNumber          int              `gorm:"index;type:integer;not null;"`
	TradeSide          types.TradeSide  `gorm:"index;type:char(4);not null;"` // buy or sell
	AvailableAssetNum  float64          `gorm:"index;type:numeric(40,15);not null;"`
	Price              float64          `gorm:"index;type:numeric(40,15);not null;"` // price in fiat
	MinAssetNum        float64          `gorm:"index;type:numeric(40,15);not null;"`
	MaxAssetNum        float64          `gorm:"index;type:numeric(40,15);not null;"`
	Updated            int64            `gorm:"autoUpdateTime:nano;not null;"`
	PaymentMethodsList pq.StringArray   `gorm:"type:text[]"`
	PaymentMethods     []*PaymentMethod `gorm:"many2many:ads_payment_methods;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;foreignKey:ID;joinForeignKey:AdvertisementID;References:ID;joinReferences:PaymentMethodID"`
	Advertiser         *Advertiser      `gorm:"references:ID"`
}

func (ads *Advertisement) BeforeCreate(tx *gorm.DB) (err error) {

	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "advertisement_id"}, {Name: "exchange_id"}},
		UpdateAll: true,
	})

	return nil
}
