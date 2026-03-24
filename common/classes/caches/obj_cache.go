package caches

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/buckets"
	"github.com/starter-go/media-pool/common/classes/objects"
)

type Cache interface {
	GetObject(id objects.ID) *CachedObject

	NewTempFile(prefix, suffix string) afs.Path
}

type CachedObject interface {
	GetID() objects.ID

	GetPath() objects.Path

	GetFile(suffix string) *CachedFile
}

type CachedFile struct {
	Object CachedObject

	Suffix string

	ID objects.ID

	Path objects.Path // the abs-path

	File afs.Path // the local-file

	Name buckets.ObjectName // the remote-name
}
