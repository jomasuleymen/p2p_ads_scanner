package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/repositories"
)

type SchemeHandler struct {
	schemeRepo *repositories.SchemeRepository
}

func NewSchemeHandler(schemeRepo *repositories.SchemeRepository) *SchemeHandler {
	return &SchemeHandler{schemeRepo: schemeRepo}
}

func (handler *SchemeHandler) BaseHandler(c *gin.Context) {
	schemeNum, _ := c.GetQuery("schemeNum")

	var schemeFilter *exchange.P2PSchemeFilter
	c.BindJSON(&schemeFilter)

	if schemeNum == "2" {
		response := handler.schemeRepo.GetTwoActionSchemes(schemeFilter)
		c.JSON(200, response)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Scheme number not supported",
	})
}
