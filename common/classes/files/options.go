package files

import (
	"os"

	"github.com/starter-go/afs"
)

func GetOptionForMakeDir() *afs.Options {
	o := &afs.Options{
		Permission: 0755,
		Flag:       os.O_CREATE,
	}
	return o
}

func GetOptionForCreateFile() *afs.Options {
	o := &afs.Options{
		Permission: 0644,
		Flag:       os.O_CREATE | os.O_WRONLY,
	}
	return o
}

func GetOptionForRewriteFile() *afs.Options {
	o := &afs.Options{
		Permission: 0644,
		Flag:       os.O_TRUNC | os.O_WRONLY,
	}
	return o
}

func GetOptionForReadFile() *afs.Options {
	o := &afs.Options{
		Permission: 0644,
		Flag:       os.O_RDONLY,
	}
	return o
}
