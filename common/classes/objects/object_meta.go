package objects

import (
	"fmt"
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
	src := inst.Meta

	h[META_ID] = inst.ID.String()
	h[META_SUM] = inst.Sum.String()
	h[META_TYPE] = inst.Type
	h[META_NAME] = inst.Name
	h[META_PATH] = inst.Path.String()
	h[META_LENGTH] = strconv.FormatInt(size, 10)

	for k, v := range src {
		if v == "" {
			continue
		}
		h[k] = v
	}

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

	dst.SetProperty("this.object", id)
	return dst
}

func LoadMetaHeaders(src properties.Table) (MetaHeaders, error) {

	if src == nil {
		return nil, fmt.Errorf("source properties.Table is nil")
	}

	gett := src.Getter()
	id := gett.GetString("this.object")
	prefix := "object." + id + "."
	headers := make(MetaHeaders)

	err := gett.Error()
	if err != nil {
		return nil, err
	}

	names := []MetaFieldName{
		META_ID,
		META_NAME,
		META_LENGTH,
		META_PATH,
		META_SUM,
		META_TYPE,
	}

	for _, shortName := range names {
		fullName := MetaFieldName(prefix) + shortName
		value := gett.GetString(string(fullName))
		headers.SetValue(shortName, value)
	}

	err = gett.Error()
	if err != nil {
		return nil, err
	}

	return headers, nil
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
	dst.Meta = headers

	return dst
}

func (headers MetaHeaders) Init() MetaHeaders {
	if headers == nil {
		headers = make(MetaHeaders)
	}
	return headers
}

func (headers MetaHeaders) SetValue(name MetaFieldName, value string) MetaHeaders {
	h := headers.Init()
	h[name] = value
	return h
}

func (headers MetaHeaders) GetValue(name MetaFieldName) string {
	h := headers
	if h == nil {
		return ""
	}
	return h[name]
}

////////////////////////////////////////////////////////////////////////////////
// EOF
