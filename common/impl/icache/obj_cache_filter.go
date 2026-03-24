package icache

import (
	"fmt"
	"io"

	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/classes/streams"
	"github.com/starter-go/vlog"
)

type ObjectCacheFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	MPCacheDir string //starter:inject("${media-pool.cache.dir}")
	FS         afs.FS //starter:inject("#")

	cacheDirPath afs.Path
}

// Put implements objects.UploadFilter.
func (inst *ObjectCacheFilterLayer) Put(o *objects.Object, next objects.UploadFilterChain) error {

	inst.innerPrepareCacheFilePath(o)

	fileData, fileMeta := inst.innerGetDataAndMetaFile(o)
	vlog.Debug("file(data) = %s", fileData.GetPath())
	vlog.Debug("file(meta) = %s", fileMeta.GetPath())

	err := inst.innerPutDataFile(o)
	if err != nil {
		return err
	}

	err = inst.innerPutMetaFile(o)
	if err != nil {
		return err
	}

	return next.Put(o)
}

// Fetch implements objects.DownloadFilter.
func (inst *ObjectCacheFilterLayer) Fetch(o *objects.Object, next objects.DownloadFilterChain) error {

	inst.innerPrepareCacheFilePath(o)

	fileData, fileMeta := inst.innerGetDataAndMetaFile(o)
	vlog.Debug("file(data) = %s", fileData.GetPath())
	vlog.Debug("file(meta) = %s", fileMeta.GetPath())

	return next.Fetch(o)
}

func (inst *ObjectCacheFilterLayer) innerPrepareCacheFilePath(o *objects.Object) {

	p1 := o.Path
	p2d := p1
	p2m := p1 + ".meta"

	o.Files.Data = &objects.CacheFile{Path: p2d}
	o.Files.Meta = &objects.CacheFile{Path: p2m}

	inst.innerMakeCacheFileComplete(o.Files.Data)
	inst.innerMakeCacheFileComplete(o.Files.Meta)
}

func (inst *ObjectCacheFilterLayer) innerMakeCacheFileComplete(cf *objects.CacheFile) {

	path := cf.Path.String()
	dir := inst.cacheDirPath
	file := dir.GetChild(path)

	cf.File = file
}

func (inst *ObjectCacheFilterLayer) innerGetDataAndMetaFile(o *objects.Object) (data, meta afs.Path) {

	d := o.Files.Data
	m := o.Files.Meta

	data = d.File
	meta = m.File
	return
}

func (inst *ObjectCacheFilterLayer) innerPutDataFile(o *objects.Object) error {

	cf := o.Files.Data
	file := cf.File

	if file.Exists() {
		return nil // skip
	}

	dir := file.GetParent()
	opt1 := streams.GetFileOptionsToMkdir()
	opt2 := streams.GetFileOptionsToCreateFile()

	if !dir.Exists() {
		err := dir.Mkdirs(opt1)
		if err != nil {
			return err
		}
	}

	src, err := o.Data.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := file.GetIO().OpenWriter(opt2)
	if err != nil {
		return err
	}

	defer dst.Close()

	cb, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	if cb != o.Size {
		return fmt.Errorf("bad size")
	}

	return nil
}

func (inst *ObjectCacheFilterLayer) innerMakeMetaFor(o *objects.Object) *objects.Meta {

	meta := new(objects.Meta)
	meta.Info = *o
	return meta
}

func (inst *ObjectCacheFilterLayer) innerPutMetaFile(o *objects.Object) error {

	cf := o.Files.Meta
	file := cf.File

	if file.Exists() {
		return nil // skip
	}

	opt2 := streams.GetFileOptionsToCreateFile()
	meta := inst.innerMakeMetaFor(o)
	text := meta.String()

	dst, err := file.GetIO().OpenWriter(opt2)
	if err != nil {
		return err
	}

	defer dst.Close()

	bin := []byte(text)
	_, err = dst.Write(bin)
	return err
}

// ListFilters implements objects.FilterRegistry.
func (inst *ObjectCacheFilterLayer) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PriorityCache,
		Label:    "ObjectCacheFilterLayer",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *ObjectCacheFilterLayer) onCreate() error {

	path := inst.MPCacheDir
	dir := inst.FS.NewPath(path)
	inst.cacheDirPath = dir

	return nil
}

func (inst *ObjectCacheFilterLayer) Life() *application.Life {
	l := new(application.Life)
	l.OnCreate = inst.onCreate
	return l
}

func (inst *ObjectCacheFilterLayer) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
