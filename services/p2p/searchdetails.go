package p2p

import (
	"fmt"

	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/samber/lo"
)

/*
TP2PSearchDetails is a global variable that contains all the information about the search data
*/
type TP2PSearchDetails struct {
	Fiats            []*models.Fiat
	Assets           []*models.Asset
	Exchanges        []*models.Exchange
	P2PExchangesData []*ExchangeP2PSearchDetails
}

func (searchData *TP2PSearchDetails) GetOrCreateState(exchangeModel *models.Exchange) *ExchangeP2PSearchDetails {
	for _, s := range searchData.P2PExchangesData {
		if s.ExchangeName == exchangeModel.Name {
			return s
		}
	}

	searchData.Exchanges = append(searchData.Exchanges, exchangeModel)

	state := NewP2PSearchState(exchangeModel.Name)
	searchData.P2PExchangesData = append(searchData.P2PExchangesData, state)

	return state
}

func (details *TP2PSearchDetails) AddFiat(fiat *models.Fiat) {
	if !lo.ContainsBy(details.Fiats, func(f *models.Fiat) bool {
		return f.Name == fiat.Name
	}) {
		details.Fiats = append(details.Fiats, fiat)
	}
}

func (details *TP2PSearchDetails) AddAsset(asset *models.Asset) {
	if !lo.ContainsBy(details.Assets, func(a *models.Asset) bool {
		return a.Name == asset.Name
	}) {
		details.Assets = append(details.Assets, asset)
	}
}

var P2PGlobalDetails = &TP2PSearchDetails{
	Fiats:            make([]*models.Fiat, 0),
	Assets:           make([]*models.Asset, 0),
	Exchanges:        make([]*models.Exchange, 0),
	P2PExchangesData: make([]*ExchangeP2PSearchDetails, 0),
}

type ExchangeP2PSearchDetails struct {
	ExchangeName  string
	Fiats         []string
	Assets        []string
	CurrencyPairs []*p2pCurrenyPair
}

type p2pCurrenyPair struct {
	Fiat  string
	Asset string
}

func NewP2PSearchState(exchangeName string) *ExchangeP2PSearchDetails {
	return &ExchangeP2PSearchDetails{
		ExchangeName:  exchangeName,
		Fiats:         make([]string, 0),
		Assets:        make([]string, 0),
		CurrencyPairs: make([]*p2pCurrenyPair, 0),
	}
}

func (state *ExchangeP2PSearchDetails) AddFiat(fiat *models.Fiat) {
	if !lo.ContainsBy(state.Fiats, func(f string) bool {
		return f == fiat.Name
	}) {
		state.Fiats = append(state.Fiats, fiat.Name)
	}
}

func (state *ExchangeP2PSearchDetails) AddAsset(asset *models.Asset) {
	if !lo.ContainsBy(state.Assets, func(a string) bool {
		return a == asset.Name
	}) {
		state.Assets = append(state.Assets, asset.Name)
	}
}

func (state *ExchangeP2PSearchDetails) AddCurrencyPairs(fiat *models.Fiat, asset *models.Asset) error {

	isCurrencyPairsExists := lo.ContainsBy(state.CurrencyPairs, func(currenyPair *p2pCurrenyPair) bool {
		return currenyPair.Fiat == fiat.Name && currenyPair.Asset == asset.Name
	})

	if isCurrencyPairsExists {
		return fmt.Errorf("currency pair %s %s is already running for %s", fiat.Name, asset.Name, state.ExchangeName)
	}

	state.AddFiat(fiat)
	state.AddAsset(asset)

	state.CurrencyPairs = append(state.CurrencyPairs, &p2pCurrenyPair{Fiat: fiat.Name, Asset: asset.Name})

	return nil
}
