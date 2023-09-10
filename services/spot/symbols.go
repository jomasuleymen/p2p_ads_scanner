package spot

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jomasuleymen/p2p_mining/exchange"

	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/jomasuleymen/p2p_mining/utils"
	"github.com/samber/lo"
)

func getBookTickerModels(exchange *exchange.Exchange, assets []*models.Asset) []*models.BookTicker {
	var bookTickers []*models.BookTicker
	var alreadyAddedSymbols []string

	allSymbols, err := fetchAllSymbols(exchange.Spot)

	if err != nil {
		fmt.Println(exchange.Name, "error while fetching spot symbols", err.Error())
		return nil
	}

	for _, asset1 := range assets {
		for _, asset2 := range assets {
			if asset1.Name == asset2.Name {
				continue
			}

			var bookTicker *models.BookTicker

			if symbol := exchange.Spot.GetSymbol(asset1.Name, asset2.Name); lo.Contains(allSymbols, symbol) {
				bookTicker = &models.BookTicker{
					ExchangeID: exchange.Model.ID,
					BaseID:     asset1.ID,
					QuoteID:    asset2.ID,
					Symbol:     symbol,
				}
			} else if symbol := exchange.Spot.GetSymbol(asset2.Name, asset1.Name); lo.Contains(allSymbols, symbol) {
				bookTicker = &models.BookTicker{
					ExchangeID: exchange.Model.ID,
					BaseID:     asset2.ID,
					QuoteID:    asset1.ID,
					Symbol:     symbol,
				}
			}

			if bookTicker == nil {
				fmt.Printf("Exchange %s doesn't support %s/%s pair\n", exchange.Name, asset1.Name, asset2.Name)
				continue
			}

			if lo.Contains(alreadyAddedSymbols, bookTicker.Symbol) {
				continue
			}

			bookTickers = append(bookTickers, bookTicker)
			alreadyAddedSymbols = append(alreadyAddedSymbols, bookTicker.Symbol)
		}
	}

	return bookTickers
}

func fetchAllSymbols(spot exchange.ISpotSymbolsFetcher) ([]string, error) {
	var tryFetchCount = 3
	var tryFetchDelay = 5 * time.Second

	var allSymbols []string

	httpClient := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 15 * time.Second,
		},
	}

	for {
		req, err := spot.SpotSymbolsHttpRequest()

		if err != nil {
			return nil, err
		}

		resp, err := httpClient.Do(req)

		if err != nil || resp.StatusCode != http.StatusOK {
			if tryFetchCount == 0 {
				return nil, err
			}

			tryFetchCount--
			time.Sleep(tryFetchDelay)
			continue
		}

		resRaw, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		allSymbols = spot.GetSpotSymbolsFromJsonRaw(resRaw)

		utils.SafeCloseRequestResponseBody(req, resp)

		break
	}

	return allSymbols, nil
}
