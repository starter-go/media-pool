package ipath

import (
	"strings"

	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/data/dxo"
)

type WebLocationMakerFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	RawBaseObjectURL string //starter:inject("${mediapool.web.objects-base-url}")

	mBaseObjectURL string
}

// Put implements objects.UploadFilter.
func (inst *WebLocationMakerFilter) Put(c *objects.IOContext, next objects.UploadFilterChain) error {
	return next.Put(c)
}

// Fetch implements objects.DownloadFilter.
func (inst *WebLocationMakerFilter) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {

	err := next.Fetch(o)
	if err != nil {
		return err
	}

	want := o.Want
	have := o.Have
	if want != nil && have != nil {
		if want.UseMeta && (want.Profile == objects.ProfileMeta) {
			url := inst.innerMakeObjectURL(have)
			have.Location = url
		}
	}

	return nil
}

func (inst *WebLocationMakerFilter) innerGetBaseObjectURL() string {

	// try get
	url := inst.mBaseObjectURL
	if url != "" {
		return url
	}

	// do load
	url = inst.RawBaseObjectURL
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	inst.mBaseObjectURL = url
	return url
}

func (inst *WebLocationMakerFilter) innerMakeObjectURL(o *objects.Object) dxo.URL {

	builder := new(strings.Builder)
	id := o.ID.String()
	name := o.Name
	base := inst.innerGetBaseObjectURL()

	if name == "" {
		name = "unnamed"
	}

	builder.WriteString(base)
	builder.WriteString(id)
	builder.WriteString("/")
	builder.WriteString(name)
	builder.WriteString("?view=meta")

	str := builder.String()
	return dxo.URL(str)
}

// ListFilters implements objects.FilterRegistry.
func (inst *WebLocationMakerFilter) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PriorityURL,
		Label:    "WebLocationMakerFilter",

		Up:   nil,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *WebLocationMakerFilter) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
