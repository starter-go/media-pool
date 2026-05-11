package icache

import (
	"encoding/hex"
	"fmt"
	"io"
	"strconv"

	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/media-pool/common/classes/files"
	"github.com/starter-go/media-pool/common/classes/layers"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/classes/streams"
	"github.com/starter-go/vlog"
)

type ObjectCacheFilterLayer struct {

	//starter:component

	_as func(objects.FilterRegistry) //starter:as(".")

	MPCacheDir string //starter:inject("${mediapool.cache.dir}")
	FS         afs.FS //starter:inject("#")

	cacheDirPath afs.Path
}

// Put implements objects.UploadFilter.
func (inst *ObjectCacheFilterLayer) Put(c *objects.IOContext, next objects.UploadFilterChain) error {

	want := c.Want

	inst.innerPrepareCacheFilePath(want)

	fileData, fileMeta := inst.innerGetDataAndMetaFile(want)
	vlog.Debug("file(data) = %s", fileData.GetPath())
	vlog.Debug("file(meta) = %s", fileMeta.GetPath())

	err := inst.innerPutDataFile(want)
	if err != nil {
		return err
	}

	err = inst.innerPutMetaFile(want)
	if err != nil {
		return err
	}

	return next.Put(c)
}

// Fetch implements objects.DownloadFilter.
func (inst *ObjectCacheFilterLayer) Fetch(c *objects.IOContext, next objects.DownloadFilterChain) error {

	want := c.Want
	countDone := 0

	inst.innerPrepareCacheFilePath(want)

	if want.UseData {
		err := inst.innerDoFetchData(c, next)
		if err != nil {
			return err
		}
		countDone++
	}

	if want.UseMeta {
		err := inst.innerDoFetchMeta(c, next)
		if err != nil {
			return err
		}
		countDone++
	}

	if want.UseThumb {
		err := inst.innerDoFetchThumb(c, next)
		if err != nil {
			return err
		}
		countDone++
	}

	if countDone > 0 {
		return nil
	}

	return inst.innerDoFetchMeta(c, next)
}

func (inst *ObjectCacheFilterLayer) innerDoFetchData(c *objects.IOContext, next objects.DownloadFilterChain) error {

	want := c.Want
	cf := want.Files.Data
	file := cf.File

	have := c.Have
	if have == nil {
		have = new(objects.Object)
		c.Have = have
	}

	if !file.Exists() {
		err := next.Fetch(c)
		if err != nil {
			return err
		}
	}

	if file.Exists() {
		have.Files.Data = cf
	}

	return nil
}

func (inst *ObjectCacheFilterLayer) innerDoFetchMeta(c *objects.IOContext, next objects.DownloadFilterChain) error {

	want := c.Want
	cf := want.Files.Meta
	file := cf.File

	if !file.Exists() {
		err := next.Fetch(c)
		if err != nil {
			return err
		}
	}

	if !file.Exists() {
		id := want.ID
		return fmt.Errorf("no wanted object-meta: object.id=%s", id)
	}

	have := c.Have
	if have == nil {
		have = new(objects.Object)
		c.Have = have
	}

	err := inst.innerLoadMetaFromCachedFile(cf, have)
	if err != nil {
		return err
	}

	have.Context = want.Context
	have.Files.Meta = want.Files.Meta
	have.Profile = want.Profile

	c.Have = have
	return nil
}

func (inst *ObjectCacheFilterLayer) innerLoadMetaFromCachedFile(src *objects.CacheFile, dst *objects.Object) error {

	file := src.File
	txt, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	pt, err := properties.Parse(txt, nil)
	if err != nil {
		return err
	}

	headers, err := objects.LoadMetaHeaders(pt)
	if err != nil {
		return err
	}

	strPath := headers.GetValue(objects.META_PATH)
	strSize := headers.GetValue(objects.META_LENGTH)
	strSum := headers.GetValue(objects.META_SUM)
	strID := headers.GetValue(objects.META_ID)

	nSize, _ := strconv.ParseInt(strSize, 10, 64)
	hexSum, _ := hex.DecodeString(strSum)

	if len(hexSum) == 32 {
		dst.Sum = objects.Sum(hexSum)
	}

	dst.ID = objects.ID(strID)
	dst.Name = headers.GetValue(objects.META_NAME)
	dst.Type = headers.GetValue(objects.META_TYPE)
	dst.Path = objects.Path(strPath)
	dst.Size = nSize
	dst.Meta = headers

	return nil
}

func (inst *ObjectCacheFilterLayer) innerDoFetchThumb(c *objects.IOContext, next objects.DownloadFilterChain) error {

	return fmt.Errorf("innerDoFetchThumb: no impl")
}

func (inst *ObjectCacheFilterLayer) innerPrepareCacheFilePath(o *objects.Object) {

	p1 := o.Path
	p2d := p1
	p2m := p1 + ".meta"

	o.Files.Data = &objects.CacheFile{Path: p2d}
	o.Files.Meta = &objects.CacheFile{Path: p2m}

	if o.UseThumb {

		// todo ...

		// sizeList := []int{32, 64, 128, 256, 512, 1024}
		// cfileMap := make(map[int]*objects.CacheFile)

		// o.Files.Thumbnail32 = cfileMap[32]
		// o.Files.Thumbnail64 = cfileMap[64]
		// o.Files.Thumbnail128 = cfileMap[128]
		// o.Files.Thumbnail256 = cfileMap[256]
		// o.Files.Thumbnail512 = cfileMap[512]
		// o.Files.Thumbnail1024 = cfileMap[1024]

	}

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

	vlog.Info("local-cache-dir = %s", dir.GetPath())

	if !dir.Exists() {
		err := dir.Mkdirs(files.GetOptionForMakeDir())
		if err != nil {
			return err
		}
	}

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
