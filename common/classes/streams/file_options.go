package streams

import (
	"io/fs"
	"os"

	"github.com/starter-go/afs"
)

func GetFileOptionsToMkdir() *afs.Options {

	const perm = (fs.ModePerm & 0755)
	const fl = os.O_RDWR | os.O_CREATE

	return &afs.Options{
		Permission: perm,
		Flag:       int(fl),
	}
}

func GetFileOptionsToCreateFile() *afs.Options {

	const perm = (fs.ModePerm & 0644)
	const fl = os.O_CREATE | os.O_WRONLY

	return &afs.Options{
		Permission: perm,
		Flag:       int(fl),
	}
}

func GetFileOptionsToRewriteFile() *afs.Options {

	const perm = (fs.ModePerm & 0644)
	const fl = os.O_TRUNC | os.O_WRONLY

	return &afs.Options{
		Permission: perm,
		Flag:       int(fl),
	}
}
