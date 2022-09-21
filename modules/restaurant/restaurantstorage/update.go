package restaurantstorage

import (
	"context"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	data.PrepareForUpdate()
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
