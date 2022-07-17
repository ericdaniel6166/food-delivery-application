package common

import (
	"errors"
	"strings"
)

type Order struct {
	OrderField string `json:"order_field" form:"order_field"`
	SortType   string `json:"sort_type" form:"sort_type"`
	//	Support cursor with UID
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (o *Order) Fulfill() {
	if o.OrderField == "" {
		o.OrderField = "id"

	}

	if o.SortType == "" {

		o.SortType = DESC

	}

}

func (o Order) Validate() error {
	o.SortType = strings.TrimSpace(o.SortType)

	if o.SortType != ASC && o.SortType != DESC {
		return errors.New("sortType should be asc or desc")
	}

	return nil

}
