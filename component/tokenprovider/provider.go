package tokenprovider

import (
	"errors"
	"food-delivery-application/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("tokenprovider not found"),
		"tokenprovider not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the tokenprovider"),
		"error encoding the tokenprovider",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(errors.New("invalid tokenprovider provided"),
		"invalid tokenprovider provided",
		"ErrInvalidToken",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
