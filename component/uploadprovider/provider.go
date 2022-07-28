package uploadprovider

import (
	"context"
	"food-delivery-application/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
