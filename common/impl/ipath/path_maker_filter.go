package ipath

import (
	"strings"

	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type PathMakerFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

}

// Put implements objects.UploadFilter.
func (inst *PathMakerFilterLayer) Put(o *objects.Object, next objects.UploadFilterChain) error {

	// make (id & path)

	id := inst.innerMakeID(o)
	o.ID = id

	path := inst.innerMakePath(o)
	o.Path = path

	return next.Put(o)
}

// Fetch implements objects.DownloadFilter.
func (inst *PathMakerFilterLayer) Fetch(o *objects.Object, next objects.DownloadFilterChain) error {

	// make (path)

	path := inst.innerMakePath(o)
	o.Path = path

	return next.Fetch(o)
}

func (inst *PathMakerFilterLayer) innerMakePath(o *objects.Object) objects.Path {

	pp := inst.innerMakePurePath(o)
	path := "objects" + pp
	return objects.Path(path)
}

func (inst *PathMakerFilterLayer) innerMakePurePath(o *objects.Object) string {

	str := o.ID.String()
	b := new(strings.Builder)
	const count = 4
	const partlen = 2

	for i := 0; i < count; i++ {
		if len(str) <= partlen {
			break
		}
		p1 := str[0:partlen]
		p2 := str[partlen:]
		b.WriteRune('/')
		b.WriteString(p1)
		str = p2
	}

	b.WriteRune('/')
	b.WriteString(str)

	return b.String()
}

func (inst *PathMakerFilterLayer) innerMakeID(o *objects.Object) objects.ID {
	sum := o.Sum
	str := sum.String()
	return objects.ID(str)
}

// ListFilters implements objects.FilterRegistry.
func (inst *PathMakerFilterLayer) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PriorityPath,
		Label:    "PathMakerFilterLayer",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *PathMakerFilterLayer) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
