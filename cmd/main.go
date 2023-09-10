package main

import (
	"github.com/jomasuleymen/p2p_mining/initializers"
	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/jomasuleymen/p2p_mining/server"
	"github.com/jomasuleymen/p2p_mining/services"
	"gorm.io/gorm"
)

func main() {
	initializers.LoadEnvVars()

	db := initializers.GetDB()

	migrateModels(db)
	services.StartServices(db)
	server.StartServer(db)

}

func migrateModels(db *gorm.DB) {
	_ = db.AutoMigrate(&models.Fiat{}, &models.Asset{}, &models.Exchange{},
		&models.AdsPaymentMethods{}, &models.PaymentMethod{},
		&models.Advertisement{}, &models.Advertiser{}, &models.BookTicker{})
}
