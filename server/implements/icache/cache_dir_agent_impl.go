package icache

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/media-pool/common/classes/caches"
	"github.com/starter-go/vlog"
)

type CacheDirAgentImpl struct {

	//starter:component

	_as func(caches.CacheDirAgent) //starter:as("#")

	MPCacheDir string //starter:inject("${mediapool.cache.dir}")
	FS         afs.FS //starter:inject("#")

	cacheDir afs.Path
}

// GetCacheDir implements caches.CacheDirAgent.
func (inst *CacheDirAgentImpl) GetCacheDir() afs.Path {

	dir := inst.cacheDir
	if dir != nil {
		return dir
	}

	p := inst.MPCacheDir
	dir = inst.FS.NewPath(p)

	inst.cacheDir = dir

	return dir
}

func (inst *CacheDirAgentImpl) onCreate() error {
	cdir := inst.GetCacheDir()
	vlog.Info("media-pool: local_cache_dir = %s", cdir.GetPath())
	return inst.innerTryMkdirs(cdir)
}

func (inst *CacheDirAgentImpl) innerTryMkdirs(dir afs.Path) error {

	if dir.Exists() {
		return nil
	}

	om := new(afs.OptionsMaker)
	om.SetMode(7, 5, 5)
	opt := om.Options()

	return dir.Mkdirs(&opt)
}

func (inst *CacheDirAgentImpl) Life() *application.Life {
	l := new(application.Life)
	l.OnCreate = inst.onCreate
	return l
}

func (inst *CacheDirAgentImpl) _impl() caches.CacheDirAgent {
	return inst
}
