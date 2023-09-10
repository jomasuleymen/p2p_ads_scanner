package bybit

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

var supportedCurrencyPairs map[string][]string

var pairsSupportMutex sync.RWMutex

func isCurrencyPairSupport(fiat string, asset string) bool {
	pairsSupportMutex.Lock()
	defer pairsSupportMutex.Unlock()

	if supportedCurrencyPairs == nil {
		supportedCurrencyPairs = make(map[string][]string)
		fetchSupportedCurrencyPairs()
	}

	if supportedCurrencyPairs[fiat] == nil {
		return false
	}

	return lo.Contains(supportedCurrencyPairs[fiat], asset)
}

type configResponseBody struct {
	Result struct {
		AllSymbol []struct {
			CurrencyId string `json:"currencyId"`
			TokenId    string `json:"tokenId"`
		} `json:"allSymbol"`
	} `json:"result"`
}

func fetchSupportedCurrencyPairs() {
	configUrl := "https://api2.bybit.com/spot/api/otc/config"

	req, err := http.NewRequest(http.MethodGet, configUrl, nil)

	if err != nil {
		log.Errorf("Error while preparing request payload: %s, %s\n", exchangeName, err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Errorf("Error while getting response: %s, %s, %s\n", exchangeName, resp.Status, err.Error())
		return
	}

	resRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error while parsing response body: %s, %s\n", exchangeName, err.Error())
		return
	}

	var config configResponseBody
	var _ = json.Unmarshal(resRaw, &config)

	for _, symbol := range config.Result.AllSymbol {
		if _, ok := supportedCurrencyPairs[symbol.CurrencyId]; !ok {
			supportedCurrencyPairs[symbol.CurrencyId] = []string{}
		}

		supportedCurrencyPairs[symbol.CurrencyId] = append(supportedCurrencyPairs[symbol.CurrencyId], symbol.TokenId)
	}

}
