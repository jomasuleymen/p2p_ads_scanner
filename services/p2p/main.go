package p2p

import (
	"github.com/jomasuleymen/p2p_mining/exchange"
	interfaces "github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/jomasuleymen/p2p_mining/repositories"
	"github.com/jomasuleymen/p2p_mining/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func StartP2PService(db *gorm.DB, exchangeServices []*interfaces.Exchange, fiats []*models.Fiat, assets []*models.Asset) {

	db.Exec("TRUNCATE TABLE advertisements CASCADE")
	db.Exec("TRUNCATE TABLE book_tickers CASCADE")
	utils.Execute_sql_file(db, "sqlqueries/ads_change_trigger.sql")

	paymentMethodRepo := NewPaymentMethodRepository(db)
	commonRepo := &repositories.CommonRepository{DB: db}

	for _, exServices := range exchangeServices {
		var exchangeModel *models.Exchange = commonRepo.UpsertExchange(exServices.Name)

		for _, fiat := range fiats {
			for _, asset := range assets {
				err := RunNewAdsFetchingWorker(db, paymentMethodRepo, exServices.P2PAdsFetcher,
					fiat, asset, exchangeModel)

				if err != nil {
					log.Error(err.Error())
				}
			}
		}
	}
}

func RunNewAdsFetchingWorker(db *gorm.DB,
	paymentRepo *PaymentRepository, adsFetcher exchange.IP2PAdsFetcher,
	fiat *models.Fiat, asset *models.Asset, exchange *models.Exchange) error {

	P2PGlobalDetails.AddFiat(fiat)
	P2PGlobalDetails.AddAsset(asset)
	var exchangeState *ExchangeP2PSearchDetails = P2PGlobalDetails.GetOrCreateState(exchange)

	if err := exchangeState.AddCurrencyPairs(fiat, asset); err != nil {
		return err
	}

	var adsRepo = NewP2PRepository(db,
		exchange, fiat, asset,
		adsFetcher, paymentRepo,
	)

	return adsRepo.StartAdsFetchingWorker()
}
