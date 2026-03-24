package itempfile

import (
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type TempFileFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

}

// Put implements objects.UploadFilter.
func (inst *TempFileFilterLayer) Put(o *objects.Object, next objects.UploadFilterChain) error {

	return next.Put(o)
}

// Fetch implements objects.DownloadFilter.
func (inst *TempFileFilterLayer) Fetch(o *objects.Object, next objects.DownloadFilterChain) error {

	return next.Fetch(o)
}

// ListFilters implements objects.FilterRegistry.
func (inst *TempFileFilterLayer) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PriorityTemp,
		Label:    "TempFileFilterLayer",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *TempFileFilterLayer) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
