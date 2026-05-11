package pools

import (
	"context"

	"github.com/starter-go/buckets"
)

type BucketLoader interface {
	LoadBucket(c context.Context) (buckets.Bucket, error)
}
