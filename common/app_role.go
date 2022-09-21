package common

//go:generate go run github.com/dmarkham/enumer -type=AppRole -json -sql -transform=snake-upper
type AppRole int64

const (
	RoleUser AppRole = 1 << iota
	RoleShipper
	RoleModerator
	RoleAdmin
	RoleAnonymous AppRole = 0
)
