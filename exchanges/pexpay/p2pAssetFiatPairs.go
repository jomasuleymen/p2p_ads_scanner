package pexpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/samber/lo"
)

var checkSupportMutex sync.RWMutex

var supportedCurrencyPairs = make(map[string][]string)

func isCurrencyPairSupport(fiat string, asset string) bool {
	checkSupportMutex.Lock()
	defer checkSupportMutex.Unlock()

	if _, ok := supportedCurrencyPairs[fiat]; !ok || len(supportedCurrencyPairs[fiat]) == 0 {
		supportedCurrencyPairs[fiat] = []string{}
		fetchSupportedCurrencyPairs(fiat)
	}

	return lo.Contains(supportedCurrencyPairs[fiat], asset)
}

type configResponseBody struct {
	Data struct {
		Areas []struct {
			Area       string `json:"area"`
			TradeSides []struct {
				Assets []struct {
					Asset string `json:"asset"`
				} `json:"assets"`
			} `json:"tradeSides"`
		} `json:"areas"`
	} `json:"data"`
}

func fetchSupportedCurrencyPairs(fiat string) {
	configUrl := "https://www.pexpay.com/bapi/c2c/v1/friendly/c2c/portal/config"
	payload := []byte(fmt.Sprintf(`{"fiat":"%s"}`, fiat))

	req, err := http.NewRequest(http.MethodPost, configUrl, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Errorf("Error while preparing request payload: %s, %s, %s\n", exchangeName, fiat, err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Errorf("Error while getting response: %s, %s, %s, %s\n", exchangeName, fiat, resp.Status, err.Error())
		return
	}

	resRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error while parsing response body: %s, %s, %s\n", exchangeName, fiat, err.Error())
		return
	}

	var config configResponseBody
	var _ = json.Unmarshal(resRaw, &config)

	for _, areaData := range config.Data.Areas {
		if areaData.Area == "P2P" {
			for _, tradeSide := range areaData.TradeSides {
				for _, asset := range tradeSide.Assets {
					if !lo.Contains(supportedCurrencyPairs[fiat], asset.Asset) {
						supportedCurrencyPairs[fiat] = append(supportedCurrencyPairs[fiat], asset.Asset)
					}
				}
			}
		}
	}

}
