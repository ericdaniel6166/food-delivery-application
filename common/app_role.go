package common

type AppRole int

const (
	User AppRole = 1 << iota
	Shipper
	Admin
)
