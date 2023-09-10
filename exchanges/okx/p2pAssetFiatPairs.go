package okx

import (
	"encoding/json"
	"net/http"

	"github.com/jomasuleymen/p2p_mining/types"
)

type configResponseBody struct {
	Code int `json:"code"`
}

func isCurrencyPairSupport(fiat string, asset string) bool {
	configUrl := "https://www.okx.com/v3/c2c/tradingOrders/books"
	req, err := http.NewRequest(http.MethodGet, configUrl, nil)

	if err != nil {
		return false
	}

	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	query := req.URL.Query()
	query.Set("side", string(types.BUY))
	query.Set("quoteCurrency", fiat)
	query.Set("baseCurrency", asset)

	req.URL.RawQuery = query.Encode()

	resp, _ := http.DefaultClient.Do(req)

	var body configResponseBody
	_ = json.NewDecoder(resp.Body).Decode(&body)

	return body.Code == 0
}
