package objects

import (
	"fmt"
	"strconv"

	"github.com/starter-go/afs"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
)

////////////////////////////////////////////////////////////////////////////////

type MetaFieldName string

const (
	META_NAME   MetaFieldName = "name"
	META_LENGTH MetaFieldName = "length"
	META_SUM    MetaFieldName = "sum"
	META_TYPE   MetaFieldName = "type"
	META_ID     MetaFieldName = "id"
	META_PATH   MetaFieldName = "path"

	META_DATE       MetaFieldName = "date"
	META_CREATED_AT MetaFieldName = "created-at"
	META_UPDATED_AT MetaFieldName = "updated-at"
)

////////////////////////////////////////////////////////////////////////////////

// 以 struct 的形式表示 meta
type Meta struct {

	// headers

	Map MetaHeaders

	// fields

	ID ID

	Name string // the simple-name of file

	Sum Sum

	Path Path

	Length int64

	Type string

	CreatedAt lang.Time

	UpdatedAt lang.Time
}

// 以 map 的形式表示 meta
type MetaHeaders map[MetaFieldName]string

// 以 properties 的形式表示 meta
type MetaProperties = properties.Table

////////////////////////////////////////////////////////////////////////////////

type MetaContext struct {
	File afs.Path

	Meta *Meta

	Map MetaHeaders

	Properties MetaProperties
}

////////////////////////////////////////////////////////////////////////////////

func (inst *Meta) Headers() MetaHeaders {

	h := make(MetaHeaders)
	size := inst.Length
	src := inst.Map
	date1 := inst.CreatedAt
	date2 := inst.UpdatedAt

	h[META_ID] = inst.ID.String()
	h[META_SUM] = inst.Sum.String()
	h[META_TYPE] = inst.Type
	h[META_NAME] = inst.Name
	h[META_PATH] = inst.Path.String()
	h[META_LENGTH] = strconv.FormatInt(size, 10)

	h[META_DATE] = strconv.FormatInt(date2.Int(), 10)
	h[META_CREATED_AT] = strconv.FormatInt(date1.Int(), 10)
	h[META_UPDATED_AT] = strconv.FormatInt(date2.Int(), 10)

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

		META_DATE,
		META_CREATED_AT,
		META_UPDATED_AT,
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
	strDate1 := src[META_CREATED_AT]
	strDate2 := src[META_UPDATED_AT]

	size, _ := strconv.ParseInt(strLen, 10, 64)
	nDate1, _ := strconv.ParseInt(strDate1, 10, 64)
	nDate2, _ := strconv.ParseInt(strDate2, 10, 64)

	hexSum := lang.Hex(strSum)
	binSum := hexSum.Bytes()

	dst.Map = headers

	dst.Sum = Sum(binSum)
	dst.ID = ID(strID)
	dst.Name = strName
	dst.Type = strType
	dst.Length = size
	dst.Path = Path(strPath)
	dst.CreatedAt = lang.Time(nDate1)
	dst.UpdatedAt = lang.Time(nDate2)

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
