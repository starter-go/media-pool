package istorage

import (
	"context"
	"fmt"
	"io"

	"github.com/starter-go/afs"
	"github.com/starter-go/buckets"
	"github.com/starter-go/media-pool/common/classes/files"
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
	want := o.Want
	file1data := want.Files.Data
	file2meta := want.Files.Meta
	ctype := want.Type

	if want.UseData {
		err = inst.innerPutFile(want, ctype, file1data)
		if err != nil {
			return err
		}
	}

	if want.UseMeta {
		err = inst.innerPutFile(want, "text/plain", file2meta)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *ObjectStoragePoolFilter) innerPutFile(o *objects.Object, ctype string, cf *objects.CacheFile) error {

	ctx := o.Context
	bucket, err := inst.BH.GetBucket(ctx)
	if err != nil {
		return err
	}

	o2 := new(buckets.Object)
	o2.Context = ctx
	o2.Name = buckets.ObjectName(cf.Path)
	o2.Type = buckets.ContentType(ctype)
	o2.Meta = map[string]string{
		"content-type": ctype,
		"obj-id":       o.ID.String(),
	}

	exi, err := bucket.Exists(o2)
	if err != nil {
		return err
	}
	if exi {
		return nil // skip
	}

	// src
	file := cf.File
	src, err := file.GetIO().OpenReader(nil)
	if err != nil {
		return err
	}
	defer src.Close()

	// put
	o2.Data = src
	o3, err := bucket.Put(o2)
	if err != nil {
		return err
	}

	vlog.Debug("bucket: put object at [%s]", o3.Name)

	return nil
}

func (inst *ObjectStoragePoolFilter) innerFetchFile(cc context.Context, want, have *objects.Object, cf *objects.CacheFile) error {

	bucket, err := inst.BH.GetBucket(cc)
	if err != nil {
		return err
	}

	o2 := new(buckets.Object)
	o2.Context = cc
	o2.Name = buckets.ObjectName(cf.Path)

	o3, err := bucket.Fetch(o2)
	if err != nil {
		return err
	}

	// src

	srcReader := o3.Data
	if srcReader == nil {
		return fmt.Errorf("data is nil")
	}
	defer srcReader.Close()

	// dst

	var dstOptions *afs.Options
	dstFile := cf.File
	dstIO := dstFile.GetIO()

	if dstFile.Exists() {
		dstOptions = files.GetOptionForRewriteFile()
	} else {
		dstOptions = files.GetOptionForCreateFile()
	}

	dstWriter, err := dstIO.OpenWriter(dstOptions)
	if err != nil {
		return err
	}
	defer dstWriter.Close()

	// pump

	count, err := io.Copy(dstWriter, srcReader)
	if err != nil {
		return err
	}

	// make source

	resultData, err := objects.MakeSourceWithCacheFile(cf)
	if err != nil {
		return err
	}

	vlog.Info("ObjectStoragePoolFilter: fetch %d byte(s) @path:%s", count, cf.Path)

	// keep result
	have.Name = want.Name
	have.Path = want.Path
	have.Data = resultData
	return nil
}

// Fetch implements objects.DownloadFilter.
func (inst *ObjectStoragePoolFilter) Fetch(o *objects.IOContext, next objects.DownloadFilterChain) error {

	ctx := o.CC
	want := o.Want
	have := &objects.Object{}

	file1data := want.Files.Data
	file2meta := want.Files.Meta

	// meta
	if want.UseMeta {
		err := inst.innerFetchFile(ctx, want, have, file2meta)
		if err != nil {
			return err
		}
	}

	// data
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
