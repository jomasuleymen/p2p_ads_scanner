package spot

import (
	"fmt"
	"io"
	"net/http"
	"time"

	interfaces "github.com/jomasuleymen/p2p_mining/exchange"

	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/jomasuleymen/p2p_mining/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func StartSpotService(db *gorm.DB, exchanges []*interfaces.Exchange, assets []*models.Asset) {
	for _, exchange := range exchanges {
		go fetchExchangeSpotPrices(db, exchange, assets)
	}
}

func fetchExchangeSpotPrices(db *gorm.DB, exchange *interfaces.Exchange, assets []*models.Asset) {
	if exchange.Spot == nil {
		fmt.Println(exchange.Name, "doesn't support spot trading")
		return
	}

	var tryFetchDelay = 5 * time.Second

	httpClient := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 15 * time.Second,
		},
	}

	var availableBookTickers []*models.BookTicker = getBookTickerModels(exchange, assets)

	for {
		req, _ := exchange.Spot.BookTickersHttpRequest()
		resp, err := httpClient.Do(req)

		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Printf(exchange.Name, "failed while fetching book tickers: %v, status code: %d", err, resp.StatusCode)

			time.Sleep(tryFetchDelay)
			continue
		}

		resRaw, err := io.ReadAll(resp.Body)

		utils.SafeCloseRequestResponseBody(req, resp)

		if err != nil {
			fmt.Println(err)
			break
		}

		var allBookTickers []interfaces.IBookTicker = exchange.Spot.GetBookTickersFromJsonRaw(resRaw)

		for _, newBookTicker := range allBookTickers {
			for _, updateBookTicker := range availableBookTickers {
				if newBookTicker.GetSymbol() == updateBookTicker.Symbol {
					updateBookTicker.BidPrice = newBookTicker.GetBidPrice()
					updateBookTicker.AskPrice = newBookTicker.GetAskPrice()
				}
			}
		}

		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "exchange_id"}, {Name: "base_id"}, {Name: "quote_id"}},
			UpdateAll: true,
		}).Create(&availableBookTickers)

		time.Sleep(tryFetchDelay)
	}

}
