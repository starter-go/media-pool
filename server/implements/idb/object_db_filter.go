package idb

import (
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/mediafiles"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type ObjectDatabaseFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	Service mediafiles.Service //starter:inject("#")

}

// Fetch implements objects.DownloadFilter.
func (inst *ObjectDatabaseFilter) Fetch(o *objects.IOC, next objects.DownloadFilterChain) error {

	return next.Fetch(o)
}

// Put implements objects.UploadFilter.
func (inst *ObjectDatabaseFilter) Put(o *objects.IOC, next objects.UploadFilterChain) error {

	err := next.Put(o)
	if err != nil {
		return err
	}

	err = inst.innerTryInsertItemToDB(o)
	if err != nil {
		// return err
		inst.innerOnError(err)
	}

	return nil
}

func (inst *ObjectDatabaseFilter) innerOnError(err error) {

}

// ListFilters implements objects.FilterRegistry.
func (inst *ObjectDatabaseFilter) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Label:    "ObjectDatabaseFilter",
		Priority: layers.PriorityDB,
		Class:    "server",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *ObjectDatabaseFilter) innerTryInsertItemToDB(o *objects.IOC) error {

	ctx := o.CC
	ser := inst.Service

	want := o.Want
	meta := &want.Meta
	item := new(mediafiles.DTO)

	// item.CreatedAt = meta.CreatedAt
	// item.UpdatedAt = meta.UpdatedAt

	item.ContentLength = meta.Length
	item.ContentType = meta.Type
	item.Path = meta.Path
	item.SimpleName = meta.Name
	item.Sum = meta.Sum.Hex()

	_, err := ser.Insert(ctx, item)

	return err
}

func (inst *ObjectDatabaseFilter) _impl() (objects.FilterRegistry, objects.DownloadFilter, objects.UploadFilter) {
	return inst, inst, inst
}
