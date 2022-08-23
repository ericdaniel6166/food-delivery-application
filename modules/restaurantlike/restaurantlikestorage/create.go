package restaurantlikestorage

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurantlike/restaurantlikemodel"
	"time"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	now := time.Now().UTC()
	data.CreatedAt = &now
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
