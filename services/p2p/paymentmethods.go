package p2p

import (
	"regexp"
	"strings"
	"sync"

	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/mehanizm/iuliia-go"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

var paymentMethodCache sync.Map
var paymentMethodIndexes sync.Map
var paymentNameRegex, _ = regexp.Compile(`[ !"#$%&'()*+,\-./:;<=>?@[\\\]^_{|}~]|(bank)|(money)`)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) *PaymentRepository {
	repo := &PaymentRepository{
		db: db,
	}

	/* Load all payment methods form database and save to in memory cache */
	var paymentMethods []*models.PaymentMethod

	repo.db.Find(&paymentMethods)

	for _, paymentMethod := range paymentMethods {
		paymentMethodCache.Store(paymentMethod.Index, paymentMethod)
	}

	return repo
}

func (repo *PaymentRepository) getPaymentMethodIndex(paymentMethod string) string {
	if index, ok := paymentMethodIndexes.Load(paymentMethod); ok {
		return index.(string)
	}

	translatedPaymentMethod := iuliia.Wikipedia.Translate(paymentMethod)
	lowerName := strings.ToLower(translatedPaymentMethod)
	withoutPunctuations := paymentNameRegex.ReplaceAllString(lowerName, "")

	paymentMethodIndexes.Store(paymentMethod, withoutPunctuations)

	return withoutPunctuations
}

func (repo *PaymentRepository) getOrCreatePaymentMethod(paymentMethod string) *models.PaymentMethod {
	paymentMethodIndex := repo.getPaymentMethodIndex(paymentMethod)

	if paymentMethodData, ok := paymentMethodCache.Load(paymentMethodIndex); ok {
		return paymentMethodData.(*models.PaymentMethod)
	}

	paymentMethodData := &models.PaymentMethod{
		Index: paymentMethodIndex,
		Name:  paymentMethod,
	}

	repo.db.Create(paymentMethodData)

	paymentMethodCache.Store(paymentMethodIndex, paymentMethodData)

	return paymentMethodData
}

func (repo *PaymentRepository) GetPaymentMethods(paymentMethodNames []string) []*models.PaymentMethod {
	var paymentMethods []*models.PaymentMethod
	var fetchedIds []int

	for _, paymentMethod := range paymentMethodNames {
		paymentEntity := repo.getOrCreatePaymentMethod(paymentMethod)

		if lo.Contains(fetchedIds, paymentEntity.ID) {
			continue
		}

		paymentMethods = append(paymentMethods, paymentEntity)
		fetchedIds = append(fetchedIds, paymentEntity.ID)
	}

	return paymentMethods
}
