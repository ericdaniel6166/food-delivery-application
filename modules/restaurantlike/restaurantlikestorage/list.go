package restaurantlikestorage

import (
	"context"
	"fmt"
	"food-delivery-application/common"
	"food-delivery-application/modules/restaurantlike/restaurantlikemodel"
	"github.com/btcsuite/btcd/btcutil/base58"
	"time"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	var listLike []restaurantlikemodel.LikedCount

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}

func (s *sqlStore) GetUsersLikeRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	order *common.Order,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var restaurantLikes []restaurantlikemodel.Like

	db := s.db

	db = db.Table(restaurantlikemodel.Like{}.TableName()).Where(conditions)
	if filter != nil {
		if filter.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", filter.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(common.TimeFormat, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format(common.TimeFormat))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&restaurantLikes).Error; err != nil {
		return nil, common.ErrDB(err)

	}

	usersLike := make([]common.SimpleUser, len(restaurantLikes))

	for i, restaurantLike := range restaurantLikes {
		userLike := restaurantLikes[i].User
		if userLike != nil {
			userLike.CreatedAt = restaurantLike.CreatedAt
			userLike.UpdatedAt = nil
			usersLike[i] = *userLike
		}
		if i == len(restaurantLikes)-1 {
			sprintf := fmt.Sprintf("%v", restaurantLike.CreatedAt.Format(common.TimeFormat))
			bytes := []byte(sprintf)
			paging.NextCursor = base58.Encode(bytes)
		}
	}

	return usersLike, nil
}
