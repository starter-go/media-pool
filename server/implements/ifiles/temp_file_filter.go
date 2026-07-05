package ifiles

import (
	"fmt"
	"io"
	"sync"

	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/common/classes/caches"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/vlog"
)

type TempFileFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	CDAgent caches.CacheDirAgent //starter:inject("#")

	tempCount int
	tempTime0 lang.Time
	tempMutex sync.Mutex
}

// Put implements objects.UploadFilter.
func (inst *TempFileFilterLayer) Put(o *objects.IOContext, next objects.UploadFilterChain) error {

	want := o.Want
	tmp := want.TempFile

	if want.UseTemp && (tmp == nil) {
		tmp = inst.innerMakeNextTempFilePath()
		want.TempFile = tmp
	}

	defer inst.innerClearTempFile(tmp)

	err := inst.innerWriteRawDataToTempFile(o, tmp)
	if err != nil {
		return err
	}

	return next.Put(o)
}

// Fetch implements objects.DownloadFilter.
func (inst *TempFileFilterLayer) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {

	// fetch 时不需要临时文件

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

func (inst *TempFileFilterLayer) innerMakeNextTempFilePath() afs.Path {
	suffix := inst.innerMakeNextTempFileNameSuffix()
	cdir := inst.CDAgent.GetCacheDir()
	return cdir.GetChild("tmp/tmp-" + suffix)
}

func (inst *TempFileFilterLayer) innerGetTime0() lang.Time {
	t0 := inst.tempTime0
	if t0 == 0 {
		t0 = lang.Now()
		inst.tempTime0 = t0
	}
	return t0
}

func (inst *TempFileFilterLayer) innerWriteRawDataToTempFile(o *objects.IOContext, tmp afs.Path) error {

	want := o.Want
	data := want.Data
	om := new(afs.OptionsMaker)

	om.Create().WriteOnly()
	om.SetMode(6, 4, 4)
	opt := om.Options()

	inst.innerMakeDirForFile(tmp)

	src, err := data.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := tmp.GetIO().OpenWriter(&opt)
	if err != nil {
		return err
	}
	defer dst.Close()

	c1 := want.Meta.Length
	c2, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	if c1 != c2 {
		vlog.Warn("bad file length, [want:%d have:%d]", c1, c2)
	}

	return nil
}

func (inst *TempFileFilterLayer) innerMakeDirForFile(file afs.Path) error {

	dir := file.GetParent()
	if dir.Exists() {
		return nil
	}

	om := new(afs.OptionsMaker)
	om.SetMode(7, 5, 5)
	opt := om.Options()

	return dir.Mkdirs(&opt)
}

func (inst *TempFileFilterLayer) innerClearTempFile(tmp afs.Path) error {
	if tmp == nil {
		return nil
	}
	if tmp.IsFile() {
		return tmp.Delete()
	}
	return nil
}

func (inst *TempFileFilterLayer) innerMakeNextTempFileNameSuffix() string {

	mu := &inst.tempMutex
	mu.Lock()
	defer mu.Unlock()

	inst.tempCount++

	count := inst.tempCount
	t1 := lang.Now()
	t0 := inst.innerGetTime0()
	const f = "-%d-%d-%d.tmp~"

	return fmt.Sprintf(f, t0, count, t1)
}

func (inst *TempFileFilterLayer) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
