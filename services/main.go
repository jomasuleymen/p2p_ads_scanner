package services

import (
	"github.com/jomasuleymen/p2p_mining/exchanges"
	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/jomasuleymen/p2p_mining/repositories"
	"github.com/jomasuleymen/p2p_mining/services/p2p"
	"github.com/jomasuleymen/p2p_mining/services/spot"
	"gorm.io/gorm"
)

var exchangeInstances = exchanges.AllExchanges

var fiats = []string{"KZT"}
var assets = []string{"USDT", "BTC", "ETH"}

func StartServices(db *gorm.DB) {

	commonRepo := repositories.NewCommonRepository(db)

	var fiatModels []*models.Fiat
	for _, fiatName := range fiats {
		fiatModels = append(fiatModels, commonRepo.UpsertFiat(fiatName))
	}

	var assetModels []*models.Asset
	for _, assetName := range assets {
		assetModels = append(assetModels, commonRepo.UpsertAsset(assetName))
	}

	for _, exchange := range exchangeInstances {
		exchange.Model = commonRepo.UpsertExchange(exchange.Name)
	}

	go spot.StartSpotService(db.Session(&gorm.Session{}), exchangeInstances, assetModels)
	go p2p.StartP2PService(db.Session(&gorm.Session{}), exchangeInstances, fiatModels, assetModels)
}
