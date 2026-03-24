package ipools

import (
	"context"

	"github.com/starter-go/buckets"
	"github.com/starter-go/media-pool/common/classes/pools"
)

type DefaultPoolService struct {

	//starter:component

	_as func(pools.Service) //starter:as("#")

	BH pools.BucketHolder //starter:inject("#")
}

// GetBucket implements pools.Service.
func (inst *DefaultPoolService) GetBucket(c context.Context) (buckets.Bucket, error) {

	return inst.BH.GetBucket(c)
}

// GetConfig implements pools.Service.
func (inst *DefaultPoolService) GetConfig(c context.Context) (*pools.Configuration, error) {
	panic("unimplemented")
}

func (inst *DefaultPoolService) _impl() pools.Service {
	return inst
}
