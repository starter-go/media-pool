package imeta

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type MetaFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

}

// Fetch implements objects.DownloadFilter.
func (inst *MetaFilter) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {

	err := next.Fetch(o)
	if err != nil {
		return err
	}

	have := o.Have
	ls := new(innerMetaLS)
	ls.file = have.Files.Meta.File

	err = ls.load()
	if err != nil {
		return err
	}

	have.Meta = *ls.meta
	return nil
}

// Put implements objects.UploadFilter.
func (inst *MetaFilter) Put(o *objects.IOContext, next objects.UploadFilterChain) error {

	want := o.Want
	ls := new(innerMetaLS)
	ls.file = want.Files.Meta.File

	err := inst.innerPrepareMeta(want, ls)
	if err != nil {
		return err
	}

	err = ls.save()
	if err != nil {
		return err
	}

	err = next.Put(o)
	if err != nil {
		return err
	}

	return inst.innerOnPutDone(o)
}

func (inst *MetaFilter) innerOnPutDone(o *objects.IOContext) error {

	want := o.GetWant(true)
	have := o.GetHave(true)

	have.Meta = want.Meta

	return nil
}

func (inst *MetaFilter) innerPrepareMeta(o *objects.Object, ls *innerMetaLS) error {
	now := lang.Now()
	meta := &o.Meta
	meta.CreatedAt = now
	meta.UpdatedAt = now
	ls.meta = meta
	return nil
}

// ListFilters implements objects.FilterRegistry.
func (inst *MetaFilter) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Label:    "MetaFilter",
		Class:    "filter",
		Enabled:  true,
		Priority: layers.PriorityMeta,

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}

}

func (inst *MetaFilter) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
