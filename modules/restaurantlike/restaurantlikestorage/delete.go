package restaurantlikestorage

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) Delete(ctx context.Context, userId, restaurantId int) error {
	db := s.db

	if err := db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
