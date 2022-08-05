package common

import "database/sql/driver"

type AppRole int64

const (
	Anonymous AppRole = 0
)

const (
	User AppRole = 1 << iota
	Shipper
	Admin
)

func (appRole *AppRole) Scan(value interface{}) error { *appRole = AppRole(value.(int64)); return nil }

func (appRole AppRole) Value() (driver.Value, error) { return int64(appRole), nil }
