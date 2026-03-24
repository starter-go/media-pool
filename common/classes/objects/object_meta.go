package objects

import (
	"strconv"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
)

type MetaFieldName string

const (
	META_NAME   MetaFieldName = "name"
	META_LENGTH MetaFieldName = "length"
	META_SUM    MetaFieldName = "sum"
	META_TYPE   MetaFieldName = "type"
	META_ID     MetaFieldName = "id"
	META_PATH   MetaFieldName = "path"
)

////////////////////////////////////////////////////////////////////////////////

type Meta struct {
	Info
}

func (inst *Meta) Headers() MetaHeaders {

	h := make(MetaHeaders)
	size := inst.Size

	h[META_ID] = inst.ID.String()
	h[META_SUM] = inst.Sum.String()
	h[META_TYPE] = inst.Type
	h[META_NAME] = inst.Name
	h[META_PATH] = inst.Path.String()
	h[META_LENGTH] = strconv.FormatInt(size, 10)

	return h
}

func (inst *Meta) String() string {
	h := inst.Headers()
	return h.String()
}

////////////////////////////////////////////////////////////////////////////////

type MetaHeaders map[MetaFieldName]string

func (headers MetaHeaders) String() string {

	props := headers.Properties()
	return properties.Format(props, properties.FormatWithGroups)
}

func (headers MetaHeaders) Properties() properties.Table {

	dst := properties.NewTable(nil)
	src := headers
	id := headers[META_ID]

	for k, value := range src {
		name := "object." + id + "." + string(k)
		dst.SetProperty(name, value)
	}

	return dst
}

func (headers MetaHeaders) Meta() *Meta {

	src := headers
	dst := new(Meta)

	strName := src[META_NAME]
	strID := src[META_ID]
	strSum := src[META_SUM]
	strType := src[META_TYPE]
	strLen := src[META_LENGTH]
	strPath := src[META_PATH]

	size, _ := strconv.ParseInt(strLen, 10, 64)

	hexSum := lang.Hex(strSum)
	binSum := hexSum.Bytes()

	dst.Sum = Sum(binSum)
	dst.ID = ID(strID)
	dst.Name = strName
	dst.Type = strType
	dst.Size = size
	dst.Path = Path(strPath)

	return dst
}

////////////////////////////////////////////////////////////////////////////////
