package imonitor

import (
	"fmt"

	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type ParamsCheckerFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	MaxContentLength int64 //starter:inject("${mediapool.web.max-content-length}")
	MinContentLength int64 //starter:inject("${mediapool.web.min-content-length}")

}

// Put implements objects.UploadFilter.
func (inst *ParamsCheckerFilter) Put(c *objects.IOContext, next objects.UploadFilterChain) error {

	err := inst.innerCheckWantPut(c)
	if err != nil {
		return err
	}

	return next.Put(c)
}

// Fetch implements objects.DownloadFilter.
func (inst *ParamsCheckerFilter) Fetch(c *objects.IOContext, next objects.DownloadFilterChain) error {

	err := inst.innerCheckWantGet(c)
	if err != nil {
		return err
	}

	want := c.Want
	meta := want.Meta

	if meta.ID == "" {
		return fmt.Errorf("ParamsCheckerFilter: param want.id is empty")
	}
	if meta.Name == "" {
		meta.Name = "unnamed.file"
	}

	return next.Fetch(c)
}

func (inst *ParamsCheckerFilter) innerCheckWantPut(c *objects.IOContext) error {

	want := c.GetWant(true)
	max := inst.MaxContentLength
	min := inst.MinContentLength
	size := want.Meta.Length

	if size < min {
		return fmt.Errorf("content-length(%d) is out of range : min(%d)", size, min)
	}
	if size > max {
		return fmt.Errorf("content-length(%d) is out of range : max(%d)", size, max)
	}

	return inst.innerCheckWantCommon(c)
}

func (inst *ParamsCheckerFilter) innerCheckWantGet(c *objects.IOContext) error {

	return inst.innerCheckWantCommon(c)
}

func (inst *ParamsCheckerFilter) innerCheckWantCommon(c *objects.IOContext) error {

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
