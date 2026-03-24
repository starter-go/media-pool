package pools

import (
	"context"

	"github.com/starter-go/buckets"
)

type Service interface {
	GetConfig(c context.Context) (*Configuration, error)

	GetBucket(c context.Context) (buckets.Bucket, error)
}
