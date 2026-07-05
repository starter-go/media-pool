package caches

import "github.com/starter-go/afs"

type CacheDirAgent interface {
	GetCacheDir() afs.Path
}
