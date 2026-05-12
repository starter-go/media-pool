package objects

import (
	"context"

	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/common/classes/streams"
	"github.com/starter-go/media-pool/common/data/dxo"
)

type ID = dxo.ObjectID

type Sum = dxo.ObjectSum

type Path = dxo.ObjectPath

////////////////////////////////////////////////////////////////////////////////

type CacheFile struct {
	File          afs.Path
	Path          Path
	ContentType   string
	ContentLength int64
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

	ThumbnailSelected *CacheFile
}

////////////////////////////////////////////////////////////////////////////////

type Object struct {

	// fields of file

	Context context.Context

	ID ID

	Sum Sum

	Path Path

	Profile Profile

	TempFile afs.Path

	UseMeta  bool
	UseData  bool
	UseThumb bool

	ThumbSize int

	// CacheFile afs.Path

	Files CacheFileSet

	Size int64

	Name string

	Location dxo.URL

	Type string // the Content-Type

	CreatedAt lang.Time // aka.'Date'

	Meta MetaHeaders

	Data streams.Source
}

type Info = Object
