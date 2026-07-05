package ifiles

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/media-pool/common/classes/caches"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type FileSetFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	CDAgent caches.CacheDirAgent //starter:inject("#")
}

// Put implements objects.UploadFilter.
func (inst *FileSetFilter) Put(o *objects.IOContext, next objects.UploadFilterChain) error {

	err := inst.innerPrepareFiles(o)
	if err != nil {
		return err
	}

	return next.Put(o)
}

// Fetch implements objects.DownloadFilter.
func (inst *FileSetFilter) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {

	err := inst.innerPrepareFiles(o)
	if err != nil {
		return err
	}

	return next.Fetch(o)
}

func (inst *FileSetFilter) innerPrepareFiles(o *objects.IOContext) error {

	want := o.Want
	builder := new(innerFileSetFilterCacheFileBuilder)
	builder.init(o, inst)

	if want.UseData {
		want.Files.Data = builder.build("")
	}

	if want.UseMeta {
		want.Files.Meta = builder.build(".meta")
	}

	if want.UseThumb {
		want.Files.Thumbnail1024 = builder.build(".1024.thumbnail")
		want.Files.Thumbnail512 = builder.build(".512.thumbnail")
		want.Files.Thumbnail256 = builder.build(".256.thumbnail")
		want.Files.Thumbnail128 = builder.build(".128.thumbnail")
		want.Files.Thumbnail64 = builder.build(".64.thumbnail")
		want.Files.Thumbnail32 = builder.build(".32.thumbnail")
	}

	return nil
}

// ListFilters implements objects.FilterRegistry.
func (inst *FileSetFilter) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Priority: layers.PriorityFiles,
		Label:    "FileSetFilter",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *FileSetFilter) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}

////////////////////////////////////////////////////////////////////////////////

type innerFileSetFilterCacheFileBuilder struct {
	cacheDir afs.Path
	objPath  objects.Path
}

func (inst *innerFileSetFilterCacheFileBuilder) init(ioc *objects.IOC, filter *FileSetFilter) {

	want := ioc.Want
	cdir := filter.CDAgent.GetCacheDir()
	op := want.Meta.Path

	inst.objPath = op
	inst.cacheDir = cdir
}

func (inst *innerFileSetFilterCacheFileBuilder) build(suffix string) *objects.CacheFile {

	base := inst.objPath
	cf := new(objects.CacheFile)
	cf.Suffix = suffix
	inst.innerMakeCacheFileComplete(base, cf)

	return cf
}

func (inst *innerFileSetFilterCacheFileBuilder) innerMakeCacheFileComplete(base objects.Path, cf *objects.CacheFile) {

	suffix := cf.Suffix
	dir := inst.cacheDir
	path := base.String() + suffix
	file := dir.GetChild(path)

	cf.File = file
	cf.Path = objects.Path(path)
}

////////////////////////////////////////////////////////////////////////////////
