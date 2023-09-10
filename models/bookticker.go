package models

type BookTicker struct {
	ExchangeID int     `gorm:"index;uniqueIndex:idx_spot_pairs;type:smallint;not null;"`
	BaseID     int     `gorm:"index;uniqueIndex:idx_spot_pairs;type:integer;not null;"`
	QuoteID    int     `gorm:"index;uniqueIndex:idx_spot_pairs;type:integer;not null;"`
	Symbol     string  `gorm:"-"`
	BidPrice   float64 `gorm:"type:decimal(21,9);not null;"`
	AskPrice   float64 `gorm:"type:decimal(21,9);not null;"`
	Updated    int64   `gorm:"autoUpdateTime:nano;not null;"`
}
