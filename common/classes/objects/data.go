package objects

import "io"

type DataProvider interface {
	Open() (io.ReadCloser, error)
}
