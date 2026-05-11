package imonitor

import (
	"fmt"

	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type ParamsCheckerFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

}

// Put implements objects.UploadFilter.
func (inst *ParamsCheckerFilter) Put(c *objects.IOContext, next objects.UploadFilterChain) error {

	err := inst.innerCheckWant(c)
	if err != nil {
		return err
	}

	return next.Put(c)
}

// Fetch implements objects.DownloadFilter.
func (inst *ParamsCheckerFilter) Fetch(c *objects.IOContext, next objects.DownloadFilterChain) error {

	err := inst.innerCheckWant(c)
	if err != nil {
		return err
	}

	want := c.Want
	if want.ID == "" {
		return fmt.Errorf("ParamsCheckerFilter: param want.id is empty")
	}
	if want.Name == "" {
		want.Name = "unnamed.file"
	}

	return next.Fetch(c)
}

func (inst *ParamsCheckerFilter) innerCheckWant(c *objects.IOContext) error {

	if c == nil {
		return fmt.Errorf("ParamsCheckerFilter: objects.IOContext is nil")
	}

	want := c.Want
	if want == nil {
		return fmt.Errorf("ParamsCheckerFilter: objects.IOContext.Want is nil")
	}

	return nil
}

// ListFilters implements objects.FilterRegistry.
func (inst *ParamsCheckerFilter) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PriorityParamsChecker,
		Label:    "ParamsCheckerFilter",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *ParamsCheckerFilter) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
