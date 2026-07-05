package files

import (
	"fmt"

	"github.com/starter-go/afs"
)

func MakeDirsForFile(file afs.Path) error {

	if file == nil {
		return fmt.Errorf("file is nil")
	}

	dir := file.GetParent()
	if dir.Exists() {
		return nil
	}

	om := new(afs.OptionsMaker)
	om.SetMode(7, 5, 5)
	opt := om.Options()

	return dir.Mkdirs(&opt)
}
