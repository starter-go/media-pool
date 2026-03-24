package objects

import (
	"context"

	"github.com/starter-go/afs"
	"github.com/starter-go/media-pool/common/classes/streams"
	"github.com/starter-go/media-pool/common/data/dxo"
)

type ID = dxo.ObjectID

type Sum = dxo.ObjectSum

type Path = dxo.ObjectPath

////////////////////////////////////////////////////////////////////////////////

type CacheFile struct {
	File afs.Path
	Path Path
}

////////////////////////////////////////////////////////////////////////////////

type CacheFileSet struct {
	Data *CacheFile
	Meta *CacheFile
}

////////////////////////////////////////////////////////////////////////////////

type Object struct {

	// fields of file

	Context context.Context

	ID ID

	Sum Sum

	Path Path

	TempFile afs.Path

	// CacheFile afs.Path

	Files CacheFileSet

	Size int64

	Name string

	Type string // the Content-Type

	Data streams.Source
}

type Info = Object
