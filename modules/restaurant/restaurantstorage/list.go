package restaurantstorage

import (
	"context"
	"food-delivery-application/common"
	restaurantmodel2 "food-delivery-application/modules/restaurant/restaurantmodel"
	"gorm.io/gorm/clause"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel2.Filter,
	paging *common.Paging,
	order *common.Order,
	moreKeys ...string,
) ([]restaurantmodel2.Restaurant, error) {
	var result []restaurantmodel2.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])

	}

	db = db.Table(restaurantmodel2.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")
	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)

		}

	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)

	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order(clause.OrderByColumn{Column: clause.Column{Name: order.OrderField}, Desc: order.SortType == common.DESC}).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)

	}

	return result, nil
}
