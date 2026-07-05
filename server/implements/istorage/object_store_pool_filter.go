package istorage

import (
	"context"

	"github.com/starter-go/buckets"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/classes/pools"
	"github.com/starter-go/vlog"
)

type ObjectStoragePoolFilter struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	BH pools.BucketHolder //starter:inject("#")
}

// Put implements objects.UploadFilter.
func (inst *ObjectStoragePoolFilter) Put(o *objects.IOContext, next objects.UploadFilterChain) error {

	var err error

	cc := o.CC
	want := o.Want
	have := o.Have
	file1data := want.Files.Data
	file2meta := want.Files.Meta
	ctype := want.Meta.Type

	if have == nil {
		have = new(objects.Object)
		o.Have = have
	}

	// for data

	if want.UseData {
		err = inst.innerPutFile(cc, want, have, file1data, ctype)
		if err != nil {
			return err
		}
	}

	// for meta

	if want.UseMeta {
		err = inst.innerPutFile(cc, want, have, file2meta, "text/plain")
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *ObjectStoragePoolFilter) innerPutFile(cc context.Context, want, have *objects.Object, cf *objects.CacheFile, ctype string) error {

	bucket, err := inst.BH.GetBucket(cc)
	if err != nil {
		return err
	}

	fileapi := bucket.ForFiles()
	o1 := new(buckets.ObjectFile)
	file := cf.File

	o1.Context = cc
	o1.Name = buckets.ObjectName(cf.Path)
	o1.Path = file
	o1.Type = buckets.ContentType(ctype)

	has, err := bucket.Exists(&o1.Object)
	if err != nil {
		return err
	}
	if has {
		return nil // skip
	}

	o2, err := fileapi.PutFile(o1)
	if err != nil {
		return err
	}

	if vlog.IsDebugEnabled() {
		file := o2.Path
		name := o2.Name
		size := o2.Size
		vlog.Debug("ObjectStoragePoolFilter.[put size:%d name:'%s' src_file:'%s' ]", size, name, file.GetPath())
	}

	return nil
}

func (inst *ObjectStoragePoolFilter) innerFetchFile(cc context.Context, want, have *objects.Object, cf *objects.CacheFile) error {

	bucket, err := inst.BH.GetBucket(cc)
	if err != nil {
		return err
	}

	fileapi := bucket.ForFiles()
	o2 := new(buckets.ObjectFile)
	file := cf.File

	if file.Exists() {
		return nil // skip
	}

	o2.Context = cc
	o2.Name = buckets.ObjectName(cf.Path)
	o2.Path = file

	o3, err := fileapi.FetchFile(o2)
	if err != nil {
		return err
	}

	if vlog.IsDebugEnabled() {
		file := o3.Path
		name := o3.Name
		size := o3.Size
		vlog.Debug("ObjectStoragePoolFilter.[fetch size:%d name:'%s' to_file:'%s' ]", size, name, file.GetPath())
	}

	// keep result

	// have.Context = cc
	// have.ID = want.ID
	// have.Name = want.Name
	// have.Path = want.Path
	// have.Profile = want.Profile
	// have.Sum = o3.Sum

	return nil
}

// Fetch implements objects.DownloadFilter.
func (inst *ObjectStoragePoolFilter) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {

	ctx := o.CC
	want := o.Want
	have := &objects.Object{}

	file1data := want.Files.Data
	file2meta := want.Files.Meta

	// for meta

	if want.UseMeta {
		err := inst.innerFetchFile(ctx, want, have, file2meta)
		if err != nil {
			return err
		}
	}

	// for data

	if want.UseData {
		err := inst.innerFetchFile(ctx, want, have, file1data)
		if err != nil {
			return err
		}
	}

	o.Have = have
	return nil
}

// ListFilters implements objects.FilterRegistry.
func (inst *ObjectStoragePoolFilter) ListFilters() []*objects.FilterRegistration {

	r1 := &objects.FilterRegistration{

		Enabled:  true,
		Label:    "ObjectStoragePoolFilter",
		Priority: layers.PriorityStorage,
		Class:    "server",

		Up:   inst,
		Down: inst,
	}

	return []*objects.FilterRegistration{r1}
}

func (inst *ObjectStoragePoolFilter) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
