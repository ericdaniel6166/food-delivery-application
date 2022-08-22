package restaurantstorage

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurant/restaurantmodel"
	"gorm.io/gorm/clause"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	order *common.Order,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")
	if filter != nil {
		if filter.CityId > 0 {
			db = db.Where("city_id = ?", filter.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])

		//if key User in different microservice, call api
		//https://github.com/go-resty/resty
		//if moreKeys[i] == "User" {
		////s.userStore.FindUsersByIds()
		////...
		//}
	}

	if fakeCursor := paging.FakeCursor; fakeCursor != "" {
		if uid, err := common.FromBase58(fakeCursor); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order(clause.OrderByColumn{Column: clause.Column{Name: order.OrderField}, Desc: order.SortType == common.DESC}).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)

	}

	return result, nil
}
