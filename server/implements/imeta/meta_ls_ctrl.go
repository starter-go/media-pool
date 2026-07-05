package imeta

import (
	"fmt"

	"github.com/starter-go/afs"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/media-pool/common/classes/files"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type innerMetaLS struct {
	file    afs.Path
	meta    *objects.Meta
	headers objects.MetaHeaders
	ptable  properties.Table
}

func (inst *innerMetaLS) load() error {

	file := inst.file
	if file == nil {
		return fmt.Errorf("file is nil")
	}

	// if !file.IsFile() {
	// }

	txt, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	pt, err := properties.Parse(txt, nil)
	if err != nil {
		return err
	}

	hdrs, err := objects.LoadMetaHeaders(pt)
	if err != nil {
		return err
	}

	meta := hdrs.Meta()

	inst.ptable = pt
	inst.headers = hdrs
	inst.meta = meta

	return nil
}

func (inst *innerMetaLS) save() error {

	meta := inst.meta
	file := inst.file

	if meta == nil {
		return fmt.Errorf("meta is nil")
	}

	if file == nil {
		return fmt.Errorf("file is nil")
	}

	hdrs1 := meta.Headers()
	hdrs2 := inst.headers.Init()
	hdrs := make(objects.MetaHeaders)
	// mix hdrs1 & hdrs2
	inst.innerCopyValues(hdrs1, hdrs)
	inst.innerCopyValues(hdrs2, hdrs)

	pt := hdrs.Properties()
	txt := properties.Format(pt, properties.FormatWithGroups)

	inst.headers = hdrs
	inst.ptable = pt

	om := new(afs.OptionsMaker)
	om.SetMode(6, 4, 4)
	om.WriteOnly()
	if !file.Exists() {
		om.Create()
	}

	files.MakeDirsForFile(file)
	opt := om.Options()

	return file.GetIO().WriteText(txt, &opt)
}

func (inst *innerMetaLS) innerCopyValues(src, dst objects.MetaHeaders) {
	if src == nil || dst == nil {
		return
	}
	for k, v := range src {
		dst.SetValue(k, v)
	}
}

func (inst *innerMetaLS) set(name objects.MetaFieldName, value string) {
	h := inst.headers
	if h == nil {
		h = h.Init()
		inst.headers = h
	}
	h.SetValue(name, value)
}
