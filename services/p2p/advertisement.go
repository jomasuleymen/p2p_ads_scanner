package p2p

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/jomasuleymen/p2p_mining/types"
	"github.com/jomasuleymen/p2p_mining/utils"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var adsInsertMutex sync.Mutex

const (
	responseErrorWaitDelay = 10 * time.Second
	nextPageFetchingDelay  = 200 * time.Millisecond
	updateDataDelay        = 30 * time.Second
)

type P2PRepository struct {
	DB                *gorm.DB
	AdsFetcher        exchange.IP2PAdsFetcher
	AdvertiserRepo    *AdvertiserRepository
	PaymentMethodRepo *PaymentRepository
	Exchange          *models.Exchange
	Fiat              *models.Fiat
	Asset             *models.Asset
	httpClient        *http.Client
	isRunning         bool
}

func NewP2PRepository(db *gorm.DB, exchange *models.Exchange, fiat *models.Fiat, asset *models.Asset,
	adsFetcher exchange.IP2PAdsFetcher, paymentMethodRepo *PaymentRepository) *P2PRepository {

	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	return &P2PRepository{DB: db,
		AdsFetcher:        adsFetcher,
		AdvertiserRepo:    NewAdvertiserRepository(db),
		PaymentMethodRepo: paymentMethodRepo,
		Exchange:          exchange,
		Fiat:              fiat, Asset: asset,
		httpClient: httpClient,
		isRunning:  false,
	}
}

func (adsRepo *P2PRepository) StartAdsFetchingWorker() error {
	exchange, fiat, asset := adsRepo.Exchange.Name, adsRepo.Fiat.Name, adsRepo.Asset.Name

	if adsRepo.isRunning {
		return fmt.Errorf("the ads fetching is already running for %s, %s, %s", fiat, asset, exchange)
	}

	if adsRepo.AdsFetcher == nil || exchange == "" || fiat == "" || asset == "" {
		return fmt.Errorf("adsFetcher, exchange, fiat or asset is not set")
	}

	if !adsRepo.AdsFetcher.CheckIfCurrencyPairSupport(fiat, asset) {
		return fmt.Errorf("currency pair %s %s is not supported by %s", fiat, asset, exchange)
	}

	adsRepo.isRunning = true

	go adsRepo.startFetchingAds()

	return nil
}

func (adsRepo *P2PRepository) startFetchingAds() {
	exchangeName, fiatName, assetName := adsRepo.Exchange.Name, adsRepo.Fiat.Name, adsRepo.Asset.Name

	log.Infof("Start fetching %s %s %s ads", exchangeName, fiatName, assetName)

	defer func() {
		if r := recover(); r != nil {
			adsRepo.isRunning = false
			log.Errorf("Fetching %s %s %s stopped. Message: %s", exchangeName, fiatName, assetName, r)
		}
	}()

	for {
		var adsList []*models.Advertisement

		for _, tradeSide := range types.TradeSides {

			sideAdsList, err := adsRepo.fetchAdsOnSide(tradeSide)

			if err != nil {
				log.Error(err)
				break
			}

			adsList = append(adsList, sideAdsList...)
		}

		adsRepo.UpsertAdvertisements(adsList)

		time.Sleep(updateDataDelay)
	}
}

func (adsRepo *P2PRepository) fetchAdsOnSide(tradeSide types.TradeSide) ([]*models.Advertisement, error) {
	adsFetcher := adsRepo.AdsFetcher

	var page = adsFetcher.GetStartPage()

	/*
		Firstly we need to make a map of ads, because we need to check if we already have this ad in map(yes, this really happens)
		And if we have, we need to update it, not create new one
		Then get values useing lodash
	*/
	var adsMap = make(map[string]bool)
	var adsList []*models.Advertisement

	hasNextPage := true

	/* make request for each page, while hasNextPage is true */
	for hasNextPage {
		req, err := adsFetcher.PrepareHttpRequest(adsRepo.Fiat.Name, adsRepo.Asset.Name, tradeSide, page)

		if err != nil {
			return nil, err
		}

		resp, err := adsRepo.httpClient.Do(req)

		if err != nil {
			return nil, err
		}

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			time.Sleep(responseErrorWaitDelay)
			continue
		}

		resRaw, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		adsListInfo := adsFetcher.ParseAdvertisements(resRaw)
		hasNextPage = adsListInfo.HasNextPage

		for _, advertise := range adsListInfo.Advertisements {
			formattedAdv := adsRepo.GetAdvertisement(tradeSide, advertise)

			if formattedAdv != nil && !adsMap[formattedAdv.AdvertisementID] {
				adsList = append(adsList, formattedAdv)
				adsMap[formattedAdv.AdvertisementID] = true
			}
		}

		utils.SafeCloseRequestResponseBody(req, resp)

		page += 1
		time.Sleep(nextPageFetchingDelay)
	}

	for index, ad := range adsList {
		ad.RowNumber = index + 1
	}

	return adsList, nil
}

func (repo *P2PRepository) GetAdvertisement(tradeSide types.TradeSide, ads exchange.IP2PAdvertisement) *models.Advertisement {

	paymentMethods := repo.PaymentMethodRepo.GetPaymentMethods(ads.GetPaymentMethods())

	if len(paymentMethods) == 0 {
		return nil
	}

	var paymentMethodsName []string = lo.Map(paymentMethods, func(method *models.PaymentMethod, _ int) string {
		return method.Name
	})

	return &models.Advertisement{
		AdvertisementID:    ads.GetAdvertisementId(),
		ExchangeID:         repo.Exchange.ID,
		TradeSide:          tradeSide,
		FiatID:             repo.Fiat.ID,
		AssetID:            repo.Asset.ID,
		Price:              ads.GetPrice(),
		AvailableAssetNum:  ads.GetAvailableAssetNum(),
		MinAssetNum:        ads.GetMinLimit(),
		MaxAssetNum:        ads.GetMaxLimit(),
		PaymentMethodsList: paymentMethodsName,
		PaymentMethods:     paymentMethods,
		Advertiser:         repo.GetAdvertiser(ads),
	}
}

func (repo *P2PRepository) GetAdvertiser(ads exchange.IP2PAdvertisement) *models.Advertiser {
	return &models.Advertiser{
		AdvertiserID:       ads.GetAdvertiserId(),
		NickName:           ads.GetUserName(),
		MonthOrderCount:    ads.GetUserMonthOrderCount(),
		OrderClosedPercent: ads.GetUserMonthClosedPercent(),
		ExchangeID:         repo.Exchange.ID,
	}
}

func (r *P2PRepository) UpsertAdvertisements(adsList []*models.Advertisement) {
	if len(adsList) == 0 {
		return
	}

	adsInsertMutex.Lock()
	defer adsInsertMutex.Unlock()

	/* Steps of upserting data
	1) upsert advertisers
	2) deactivate all ads by exchange, asset, fiat
	3) delete all ads_ayment_methods by id of deactivated ads (because, may be there are some ads with changed payment methods, that we can't handle it)
	4) insert ads (ads_payment_methods will be inserted automatically)
	5) delete all ads that are not updated (maybe they are deleted on the exchange)
	6) upsert ads_payment_methods
	*/

	// 1) upsert advertisers
	if advertiserError := r.AdvertiserRepo.UpsertFromAdvertisements(adsList); advertiserError != nil {
		fmt.Println("error while inserting advertisers: ", advertiserError.Error())
		return
	}

	tx := r.DB.Begin()

	var deactivatedAdsIds []uint

	// 2) deactivate all ads by exchange, asset, fiat
	tx.Model(&models.Advertisement{}).
		Where(`exchange_id = ? AND asset_id = ? AND fiat_id = ?`,
			r.Exchange.ID, r.Asset.ID, r.Fiat.ID).
		Clauses(
			clause.Locking{Strength: "UPDATE"},
			clause.Returning{Columns: []clause.Column{{Name: "id"}}}).
		Pluck("id", &deactivatedAdsIds)

	// 3) delete all ads_ayment_methods by id of deactivated ads (because, may be there are some ads with changed payment methods, that we can't handle it)
	if len(deactivatedAdsIds) > 0 {
		tx.Where("advertisement_id IN ?", deactivatedAdsIds).Delete(&models.AdsPaymentMethods{})
	}

	// 4) insert ads (ads_payment_methods will be inserted automatically)
	tx.Model(&models.Advertisement{}).Omit("Advertiser", "PaymentMethods").Create(adsList)

	// 5) delete all ads that are not updated (maybe they are deleted on the exchange)
	tx.Exec(
		`DELETE FROM advertisements
		WHERE exchange_id = ? AND asset_id = ? AND fiat_id = ? AND updated < ?`,
		r.Exchange.ID, r.Asset.ID, r.Fiat.ID, adsList[0].Updated)

	// 6) upsert ads_payment_methods
	adsPaymentMethodsError := upsertAdsPaymentMethods(tx, adsList)

	if adsPaymentMethodsError != nil {
		fmt.Println("adsPaymentMethods error: ", adsPaymentMethodsError.Error())
	}

	err := tx.Commit().Error
	if err != nil {
		fmt.Println("advertisements error: ", err.Error())
		return
	}

}

func upsertAdsPaymentMethods(tx *gorm.DB, adsList []*models.Advertisement) error {
	adsPaymentMethods := make([]*models.AdsPaymentMethods, 0)

	for _, ad := range adsList {

		for _, pm := range ad.PaymentMethods {
			adsPaymentMethods = append(adsPaymentMethods, &models.AdsPaymentMethods{
				AdvertisementID: ad.ID,
				PaymentMethodID: pm.ID,
			})
		}
	}

	result := tx.Model(&models.AdsPaymentMethods{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "advertisement_id"}, {Name: "payment_method_id"}},
		UpdateAll: true,
	}).Create(adsPaymentMethods)

	return result.Error
}
