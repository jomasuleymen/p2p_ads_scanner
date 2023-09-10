package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Exchange struct {
	ID   int    `gorm:"primaryKey;type:smallserial;autoIncrement:true;"`
	Name string `gorm:"unique;index;not null;type:varchar(35)"`
}

func (exchange *Exchange) BeforeCreate(tx *gorm.DB) (err error) {

	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		UpdateAll: true,
	})

	return nil
}
