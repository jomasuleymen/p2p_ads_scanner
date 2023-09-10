package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Advertiser struct {
	ID                 int    `gorm:"primaryKey;type:bigint;autoincrement:true"`
	AdvertiserID       string `gorm:"index;uniqueIndex:exchange_advertiser_id;type:varchar(150);not null;"`
	ExchangeID         int    `gorm:"index;uniqueIndex:exchange_advertiser_id;type:smallint;not null;"`
	NickName           string `gorm:"index;type:varchar(150);not null;"`
	MonthOrderCount    int    `gorm:"index;type:integer;not null;"`
	OrderClosedPercent int    `gorm:"index;type:integer;not null;"`
}

func (advertiser *Advertiser) BeforeCreate(tx *gorm.DB) (err error) {

	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "advertiser_id"}, {Name: "exchange_id"}},
		UpdateAll: true,
	})

	return nil
}
