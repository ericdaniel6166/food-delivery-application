package common

type AppRole int

const (
	Anonymous AppRole = 0
)

const (
	User AppRole = 1 << iota
	Shipper
	Admin
)
