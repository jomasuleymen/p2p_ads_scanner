package p2p

import (
	"sync"

	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type AdvertiserRepository struct {
	mutex sync.Mutex
	DB    *gorm.DB
}

func NewAdvertiserRepository(db *gorm.DB) *AdvertiserRepository {
	return &AdvertiserRepository{
		DB: db,
	}
}

func (r *AdvertiserRepository) insertAdvertisers(advertisers []*models.Advertiser) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.DB.Create(advertisers).Error
}

func (r *AdvertiserRepository) UpsertFromAdvertisements(adsList []*models.Advertisement) error {

	/* fetch unique advertisers, because postgresql doesn't allow update
	on clause conflict more than one time */
	var advertisersById = make(map[string]*models.Advertiser)

	for _, ad := range adsList {
		advertisersById[ad.Advertiser.AdvertiserID] = ad.Advertiser
	}

	var advertisers = lo.Values(advertisersById)

	if err := r.insertAdvertisers(advertisers); err != nil {
		return err
	}

	/* update advertisements' advertiserId for association(foreign key) after upserting data */
	for _, ad := range adsList {
		ad.AdvertiserID = advertisersById[ad.Advertiser.AdvertiserID].ID
	}

	return nil
}
