package iobjects

import "github.com/starter-go/media-pool/common/classes/objects"

type ObjectServiceImpl struct {

	//starter:component

	_as func(objects.Service) //starter:as("#")
}

// Fetch implements objects.Service.
func (inst *ObjectServiceImpl) Fetch(o *objects.Info) error {
	panic("unimplemented")
}

// Put implements objects.Service.
func (inst *ObjectServiceImpl) Put(o *objects.Info) error {
	panic("unimplemented")
}

func (inst *ObjectServiceImpl) _impl() objects.Service {
	return inst
}
