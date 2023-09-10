package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jomasuleymen/p2p_mining/repositories"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	router := gin.Default()

	schemeRepository := repositories.NewSchemeRepository(db)
	schemeHandler := NewSchemeHandler(schemeRepository)

	router.GET("/api/market-information", MarketInformationHandler)
	router.POST("/api/scheme", schemeHandler.BaseHandler)
	// router.POST("/api/spreads", SpreadHandler)
	// router.POST("/api/monitoring", GetTopAdsHandler)

	router.Run("localhost:8080")
}
