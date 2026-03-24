package pools

import (
	"context"

	"github.com/starter-go/buckets"
)

type BucketHolder interface {
	GetBucket(c context.Context) (buckets.Bucket, error)
}

type BucketLoader interface {
	LoadBucket(c context.Context) (buckets.Bucket, error)
}
