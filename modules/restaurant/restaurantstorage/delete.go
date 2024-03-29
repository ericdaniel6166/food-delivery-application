package restaurantstorage

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) SoftDeleteData(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{
		"status": false,
	}).Error; err != nil {
		return common.ErrDB(err)

	}

	return nil
}
