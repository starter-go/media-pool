package icache

import (
	"fmt"

	"github.com/starter-go/afs"
	"github.com/starter-go/media-pool/common/classes/files"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type ObjectCacheFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

}

// Put implements objects.UploadFilter.
func (inst *ObjectCacheFilterLayer) Put(c *objects.IOContext, next objects.UploadFilterChain) error {

	// 这里,只处理实体数据文件

	want := c.Want
	dataCF := want.Files.Data

	srcDataFile := want.TempFile
	dstDataFile := dataCF.File

	if srcDataFile == nil || dstDataFile == nil {
		return fmt.Errorf("ObjectCacheFilterLayer.Put() : src|dst file is nil")
	}

	if !dstDataFile.Exists() {
		om := new(afs.OptionsMaker)
		om.SetMode(6, 4, 4)
		om.Create().WriteOnly()
		opt := om.Options()
		files.MakeDirsForFile(dstDataFile)
		err := srcDataFile.MoveTo(dstDataFile, &opt)
		if err != nil {
			return err
		}
	}

	return next.Put(c)
}

// Fetch implements objects.DownloadFilter.
func (inst *ObjectCacheFilterLayer) Fetch(c *objects.IOContext, next objects.DownloadFilterChain) error {

	// 这里,只处理实体数据文件

	want := c.Want
	meta := want.Files.Meta
	data := want.Files.Data

	if meta == nil || data == nil {
		return fmt.Errorf("ObjectCacheFilterLayer.Fetch() : cache-file is nil")
	}

	metaFile := meta.File
	dataFile := data.File

	if metaFile == nil || dataFile == nil {
		return fmt.Errorf("ObjectCacheFilterLayer.Fetch() : data|meta file is nil")
	}

	if metaFile.IsFile() && dataFile.IsFile() {
		// cache is ready
		c.Have = want
		return nil
	}

	return next.Fetch(c)
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

func (inst *ObjectCacheFilterLayer) _impl() (objects.FilterRegistry, objects.UploadFilter, objects.DownloadFilter) {
	return inst, inst, inst
}
