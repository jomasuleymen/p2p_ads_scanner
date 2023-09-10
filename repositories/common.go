package repositories

import (
	"github.com/jomasuleymen/p2p_mining/models"
	"gorm.io/gorm"
)

type CommonRepository struct {
	DB *gorm.DB
}

func NewCommonRepository(db *gorm.DB) *CommonRepository {
	return &CommonRepository{DB: db}
}

func (repo *CommonRepository) UpsertFiat(fiat string) *models.Fiat {
	var entity *models.Fiat = &models.Fiat{Name: fiat}
	err := repo.DB.Create(entity).Error

	if err != nil {
		panic(err)
	}

	return entity
}

func (repo *CommonRepository) UpsertAsset(asset string) *models.Asset {
	var entity *models.Asset = &models.Asset{Name: asset}
	err := repo.DB.Create(entity).Error

	if err != nil {
		panic(err)
	}

	return entity
}

func (repo *CommonRepository) UpsertExchange(exchange string) *models.Exchange {
	var entity *models.Exchange = &models.Exchange{Name: exchange}
	err := repo.DB.Create(entity).Error

	if err != nil {
		panic(err)
	}

	return entity
}
