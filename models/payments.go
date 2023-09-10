package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentMethod struct {
	ID    int    `gorm:"primaryKey;autoIncrement:true;not null;"`
	Index string `gorm:"unique;type:varchar(65);not null;"`
	Name  string `gorm:"index;type:varchar(65);not null;"`
}

func (paymentMethod *PaymentMethod) BeforeCreate(tx *gorm.DB) (err error) {

	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "index"}},
		DoUpdates: clause.AssignmentColumns([]string{"index"}),
	})

	return nil
}
