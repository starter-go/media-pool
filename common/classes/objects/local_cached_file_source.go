package objects

import (
	"fmt"
	"io"

	"github.com/starter-go/media-pool/common/classes/streams"
)

func MakeSourceWithCacheFile(cf *CacheFile) (streams.Source, error) {
	result := new(innerSourceForCacheFile)
	err := result.init(cf)
	if err != nil {
		return nil, err
	}
	return result, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerSourceForCacheFile struct {
	cf CacheFile
}

func (inst *innerSourceForCacheFile) init(cf *CacheFile) error {

	if cf == nil {
		return fmt.Errorf("innerSourceForCacheFile.init() : param is nil")
	}

	if cf.File == nil {
		return fmt.Errorf("innerSourceForCacheFile.init() : file is nil")
	}

	if cf.Path == "" {
		return fmt.Errorf("innerSourceForCacheFile.init() : path is empty")
	}

	inst.cf = *cf
	return nil
}

// Open implements streams.Source.
func (inst *innerSourceForCacheFile) Open() (io.ReadCloser, error) {

	file := inst.cf.File
	if file == nil {
		return nil, fmt.Errorf("innerSourceForCacheFile: file is nil")
	}

	if !file.Exists() {
		path := file.GetPath()
		return nil, fmt.Errorf("innerSourceForCacheFile: the file is NOT exists. path = %s", path)
	}

	fio := file.GetIO()
	return fio.OpenReader(nil)
}

func (inst *innerSourceForCacheFile) _impl() streams.Source {
	return inst
}
