package objects

import (
	"context"

	"github.com/starter-go/afs"
	"github.com/starter-go/media-pool/common/classes/streams"
	"github.com/starter-go/media-pool/common/data/dxo"
)

type ID = dxo.ObjectID

// 表示对象的sum, 固定为 sha-256 算法
type Sum = dxo.ObjectSum

type Path = dxo.ObjectPath

////////////////////////////////////////////////////////////////////////////////

type CacheFile struct {
	File   afs.Path
	Path   Path
	Suffix string
	Type   string // content-type
	Length int64  // content-length
}

////////////////////////////////////////////////////////////////////////////////

type CacheFileSet struct {
	Data *CacheFile // '{{_sum}}'

	Meta *CacheFile // '{{_sum}}.meta'

	Thumbnail32   *CacheFile // '{{_sum}}.thumb_32.jpg'
	Thumbnail64   *CacheFile // '{{_sum}}.thumb_64.jpg'
	Thumbnail128  *CacheFile // '{{_sum}}.thumb_128.jpg'
	Thumbnail256  *CacheFile // '{{_sum}}.thumb_256.jpg'
	Thumbnail512  *CacheFile // '{{_sum}}.thumb_512.jpg'
	Thumbnail1024 *CacheFile // '{{_sum}}.thumb_1024.jpg'

	Selected *CacheFile
}

////////////////////////////////////////////////////////////////////////////////

type Object struct {

	// fields of file

	Context context.Context

	Profile Profile

	ThumbSize int

	TempFile afs.Path

	UseMeta  bool
	UseData  bool
	UseTemp  bool
	UseThumb bool

	// CacheFile afs.Path

	Files CacheFileSet

	Meta Meta

	Location dxo.URL

	Data streams.Source

	// 这些字段已经废弃 , use Meta 代替

	// ID        ID
	// Name      string
	// Sum       Sum
	// Path      Path
	// Type      string // the Content-Type
	// Size      int64
	// CreatedAt lang.Time // aka.'Date'

}

type Info = Object
