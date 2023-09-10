package repositories

import (
	"fmt"

	"github.com/jomasuleymen/p2p_mining/exchange"
	"github.com/jomasuleymen/p2p_mining/models"
	"github.com/jomasuleymen/p2p_mining/types"
	"gorm.io/gorm"
)

type SchemeRepository struct {
	DB              *gorm.DB
	PageLimit       int
	compareAdsLimit int
}

func NewSchemeRepository(db *gorm.DB) *SchemeRepository {
	return &SchemeRepository{DB: db, compareAdsLimit: 10, PageLimit: 15}
}

func (repo *SchemeRepository) GetTwoActionSchemes(filter *exchange.P2PSchemeFilter) interface{} {
	db := repo.DB

	var minSpread float64 = getSpread(filter.SpreadPercent.Min)
	var maxSpread float64 = getSpread(filter.SpreadPercent.Max)

	response := []map[string]interface{}{}

	var buyOnSide types.TradeSide = types.BUY
	var sellOnSide types.TradeSide = types.SELL

	if filter.Actions.BuyAs == exchange.Maker {
		buyOnSide = types.SELL
	}

	if filter.Actions.SellAs == exchange.Maker {
		sellOnSide = types.BUY
	}

	commonData := repo.getCommonData(filter)
	buyData := repo.getData(types.BUY, buyOnSide, filter.Payments.WeBuy)
	sellData := repo.getData(types.SELL, sellOnSide, filter.Payments.WeSell)

	buyDataFiltered := repo.getFilterData("buy_data", "sell_data", types.BUY, buyOnSide, minSpread, maxSpread)
	sellDataFiltered := repo.getFilterData("sell_data", "buy_data_f", types.SELL, sellOnSide, minSpread, maxSpread)

	resultQuery := db.Select(
		"buy.id as buy_id",
		"Round((sell.price / buy.price - 1) * 100, 3) AS spread_percent",
		"sell.id as sell_id",
	).
		Table("buy_data_f buy").
		Joins("INNER JOIN sell_data_f sell ON sell.price / buy.price BETWEEN ? AND ?", minSpread, maxSpread).
		Order("spread_percent desc").
		Limit(repo.PageLimit)

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return db.Raw(`WITH common_table as (?), buy_data as (?), sell_data as (?), buy_data_f as (?), sell_data_f as (?) ?`,
			commonData, buyData, sellData, buyDataFiltered, sellDataFiltered, resultQuery)
	})

	fmt.Println(sql)

	return response
}

func (repo *SchemeRepository) getCommonData(filter *exchange.P2PSchemeFilter) *gorm.DB {
	data := repo.DB.Model(&models.Advertisement{}).
		Select("advertisements.*", "apm.*").
		Joins("INNER JOIN ads_payment_methods AS apm ON apm.advertisement_id = advertisements.id").
		Where("asset_id = ?", filter.Assets.Selected).
		Where("fiat_id = ?", filter.Fiat)

	if len(filter.Exchanges.Declined) > 0 {
		data = data.Where("advertisements.exchange_id NOT IN ?", filter.Exchanges.Declined)
	}

	if len(filter.Payments.Declined) > 0 {
		data = data.Where("payment_method_id NOT IN ?", filter.Payments.Declined)
	}

	if len(filter.User.IgnoreIds) > 0 || filter.User.MinOrder > 0 || filter.User.MinRate > 0 {
		data = data.Joins("INNER JOIN advertisers ON advertisers.id = advertisements.advertiser_id")

		if len(filter.User.IgnoreIds) > 0 {
			data = data.Where("advertisers.id NOT IN ?", filter.User.IgnoreIds)
		}

		if filter.User.MinOrder > 0 {
			data = data.Where("advertisers.month_order_count > ?", filter.User.MinOrder)
		}

		if filter.User.MinRate > 0 {
			data = data.Where("advertisers.order_closed_percent > ?", filter.User.MinRate)
		}
	}

	return data
}

func (repo *SchemeRepository) getData(tradeSide, tradeOnSide types.TradeSide, payments []int) *gorm.DB {
	data := repo.DB.Select("*").
		Table("common_table").
		Where("trade_side = ?", tradeOnSide)

	if len(payments) > 0 {
		data = data.Where("payment_method_id IN ?", payments)
	}

	// maker
	if tradeSide != tradeOnSide {

		var order string
		if tradeOnSide == types.SELL {
			order = "DESC"
		} else {
			order = "ASC"
		}

		subQuery := data.Select("*", fmt.Sprintf("ROW_NUMBER() OVER (PARTITION BY exchange_id, payment_method_id ORDER BY price %s) as ranked_order", order))
		data = repo.DB.Select("*").
			Table("(?) as mt", subQuery).
			Where("mt.ranked_order = 1")
	}

	return data
}

func (repo *SchemeRepository) getFilterData(rootTableName, filterByTable string, actSide, tradeOnSide types.TradeSide, minSpread, maxSpread float64) *gorm.DB {
	filteredData := repo.DB.Select("*").Table(rootTableName)

	if actSide == types.BUY {
		filteredData.Where("price BETWEEN (?) / ? AND (?) / ?", repo.DB.Select("Min(price)").Table(filterByTable), maxSpread,
			repo.DB.Select("Max(price)").Table(filterByTable), minSpread)
	} else {
		filteredData.Where("price BETWEEN (?) * ? AND (?) * ?", repo.DB.Select("Min(price)").Table(filterByTable), minSpread,
			repo.DB.Select("Max(price)").Table(filterByTable), maxSpread)
	}

	/* limit only for taker advertisements.
	   For maker we don't need to limit,
	   because we will compare only activities of payment methods */

	if actSide == tradeOnSide {

		var order string
		if actSide == types.SELL {
			order = "DESC"
		} else {
			order = "ASC"
		}

		filteredData = filteredData.Order(fmt.Sprintf("price %s", order)).Limit(repo.compareAdsLimit)
	}

	return filteredData
}

func getSpread(spreadPercent float64) float64 {
	return 1 + spreadPercent/100.0
}
