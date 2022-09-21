package restaurantstorage

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	data.PrepareForInsert()
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)

	}

	return nil
}
