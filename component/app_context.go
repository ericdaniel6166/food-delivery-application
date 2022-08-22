package component

import (
	"food-delivery-application/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	Version() string
	JwtExpirationInSeconds() int
}

type appCtx struct {
	db                     *gorm.DB
	uploadProvider         uploadprovider.UploadProvider
	secretKey              string
	version                string
	jwtExpirationInSeconds int
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, version string, jwtExpirationInSeconds int) *appCtx {
	return &appCtx{db: db, uploadProvider: upProvider, secretKey: secretKey, version: version, jwtExpirationInSeconds: jwtExpirationInSeconds}

}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

func (ctx *appCtx) Version() string { return ctx.version }

func (ctx *appCtx) JwtExpirationInSeconds() int { return ctx.jwtExpirationInSeconds }
