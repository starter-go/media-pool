package ipools

import (
	"context"
	"fmt"

	"github.com/starter-go/buckets"
	"github.com/starter-go/media-pool/common/classes/pools"
)

////////////////////////////////////////////////////////////////////////////////

type DefaultBucketHolder struct {

	//starter:component

	_as func(pools.BucketHolder) //starter:as("#")

	Loader pools.BucketLoader //starter:inject("#")

	bucket buckets.Bucket
}

// GetBucket implements pools.BucketHolder.
func (inst *DefaultBucketHolder) GetBucket(c context.Context) (buckets.Bucket, error) {

	b := inst.bucket
	if b != nil {
		return b, nil
	}

	b, err := inst.innerLoadBucket(c)
	if b == nil && err == nil {
		err = fmt.Errorf("DefaultBucketHolder: bucket is nil")
	}
	if err != nil {
		return nil, err
	}

	inst.bucket = b
	return b, nil
}

func (inst *DefaultBucketHolder) innerLoadBucket(c context.Context) (buckets.Bucket, error) {
	return inst.Loader.LoadBucket(c)
}

func (inst *DefaultBucketHolder) _impl() pools.BucketHolder {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type DefaultBucketLoader struct {

	//starter:component

	_as func(pools.BucketLoader) //starter:as("#")

}

// LoadBucket implements pools.BucketLoader.
func (inst *DefaultBucketLoader) LoadBucket(c context.Context) (buckets.Bucket, error) {
	panic("unimplemented")
}

func (inst *DefaultBucketLoader) _impl() pools.BucketLoader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
