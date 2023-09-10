package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Asset struct {
	ID   int    `gorm:"primaryKey;type:smallserial;autoIncrement:true;"`
	Name string `gorm:"unique;index;not null;type:varchar(8)"`
}

func (asset *Asset) BeforeCreate(tx *gorm.DB) (err error) {

	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		UpdateAll: true,
	})

	return nil
}
