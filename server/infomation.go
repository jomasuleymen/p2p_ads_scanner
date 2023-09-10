package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jomasuleymen/p2p_mining/services/p2p"
)

func MarketInformationHandler(c *gin.Context) {
	query, _ := c.GetQuery("query")

	switch query {
	case "common-details":
		c.JSON(200, p2p.P2PGlobalDetails)
	case "exchanges-search-details":
		c.JSON(200, p2p.P2PGlobalDetails.P2PExchangesData)
	// case "payment-methods":
	// 	fiat, ok := c.GetQuery("fiat")
	// 	if ok {
	// 		c.JSON(200, )
	// 	}
	default:
		SendErrorMessage(c, fmt.Sprintf(`query with value "%s" not found`, query))
	}
}
