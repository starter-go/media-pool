package ihash

import (
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type SumFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

}

// Put implements objects.UploadFilter.
func (inst *SumFilterLayer) Put(o *objects.Object, next objects.UploadFilterChain) error {

	sum, err := inst.computeSum(o)

	// sum2, err2 := inst.computeSum(o)
	// sum3, err3 := inst.computeSum(o)
	// vlog.Debug("", sum2, sum3, err2, err3)

	if err != nil {
		return err
	}
	o.Sum = sum

	return next.Put(o)
}

func (inst *SumFilterLayer) computeSum(o *objects.Object) (objects.Sum, error) {

	var sum objects.Sum

	d := o.Data
	in, err := d.Open()
	if err != nil {
		return sum, err
	}

	defer in.Close()
	out := sha256.New()

	cb, err := io.Copy(out, in)
	if err != nil {
		return sum, err
	}
	if cb != o.Size {
		return sum, fmt.Errorf("bad size")
	}

	s2 := out.Sum(nil)
	sum = objects.Sum(s2)
	return sum, nil
}

// Fetch implements objects.DownloadFilter.
func (inst *SumFilterLayer) Fetch(o *objects.Object, next objects.DownloadFilterChain) error {

	return next.Fetch(o)
}

// ListFilters implements objects.FilterRegistry.
func (inst *SumFilterLayer) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PrioritySum,
		Label:    "SumFilterLayer",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *SumFilterLayer) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
