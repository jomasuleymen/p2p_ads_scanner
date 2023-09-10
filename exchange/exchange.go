package exchange

import "github.com/jomasuleymen/p2p_mining/models"

type Exchange struct {
	Name          string
	Model         *models.Exchange
	Spot          ISpot
	P2PAdsFetcher IP2PAdsFetcher
}
