package component

import (
	"food-delivery-application/component/uploadprovider"
	"food-delivery-application/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	Version() string
	JwtExpirationInSeconds() int
	GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db                     *gorm.DB
	uploadProvider         uploadprovider.UploadProvider
	secretKey              string
	version                string
	jwtExpirationInSeconds int
	pb                     pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, version string,
	jwtExpirationInSeconds int, pb pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, uploadProvider: upProvider, secretKey: secretKey, version: version,
		jwtExpirationInSeconds: jwtExpirationInSeconds, pb: pb}

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

func (ctx *appCtx) GetPubsub() pubsub.Pubsub { return ctx.pb }
