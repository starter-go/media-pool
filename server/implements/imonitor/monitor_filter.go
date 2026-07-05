package imonitor

import (
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/vlog"
)

type MonitorFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

}

// Put implements objects.UploadFilter.
func (inst *MonitorFilterLayer) Put(o *objects.IOContext, next objects.UploadFilterChain) error {

	err := next.Put(o)
	if err == nil {
		inst.log(o.Want, "put:")
	}
	return err
}

// Fetch implements objects.DownloadFilter.
func (inst *MonitorFilterLayer) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {

	// err := next.Fetch(o)
	// if err == nil {
	// 	inst.log(o, "fetch:")
	// }

	return next.Fetch(o)
}

func (inst *MonitorFilterLayer) log(o *objects.Object, title string) {

	if o == nil {
		return
	}

	m := &o.Meta

	sum := m.Sum.String()
	tp := m.Type
	name := m.Name
	id := m.ID.String()
	path := m.Path.String()
	size := m.Length

	vlog.Debug(title+"[object id:%s sum:%s size:%d type:%s path:%s name:%s]", id, sum, size, tp, path, name)
}

// ListFilters implements objects.FilterRegistry.
func (inst *MonitorFilterLayer) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PriorityMonitor,
		Label:    "MonitorFilterLayer",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *MonitorFilterLayer) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
