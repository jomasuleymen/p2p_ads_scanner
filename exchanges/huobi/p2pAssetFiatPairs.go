package huobi

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/samber/lo"
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
	Data []struct {
		CryptoAsset struct {
			Name string `json:"name"`
		} `json:"cryptoAsset"`
		QuoteAsset []struct {
			Name string `json:"name"`
		} `json:"quoteAsset"`
	} `json:"data"`
}

func fetchSupportedCurrencyPairs() {
	configUrl := "https://otc-akm.huobi.com/v1/trade/fast/config/list?side=buy&tradeMode=c2c_simple"

	req, err := http.NewRequest(http.MethodGet, configUrl, nil)

	if err != nil {
		log.Errorf("Error while preparing request payload: %s, %s\n", exchangeName, err.Error())
		return
	}

	req.Header.Set("client-type", "web")

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

	for _, currencyPair := range config.Data {
		cryptoAsset := currencyPair.CryptoAsset.Name
		for _, quoteAsset := range currencyPair.QuoteAsset {
			if supportedCurrencyPairs[quoteAsset.Name] == nil {
				supportedCurrencyPairs[quoteAsset.Name] = make([]string, 0)
			}

			supportedCurrencyPairs[quoteAsset.Name] = append(supportedCurrencyPairs[quoteAsset.Name], cryptoAsset)
		}
	}

}
