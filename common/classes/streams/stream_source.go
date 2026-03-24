package streams

import (
	"io"
	"mime/multipart"
)

type Source interface {
	Open() (io.ReadCloser, error)
}

////////////////////////////////////////////////////////////////////////////////

func NewSourceForMultipart(fh *multipart.FileHeader) Source {

	src := new(innerMultipartSource)
	src.header = fh
	return src
}

type innerMultipartSource struct {
	header *multipart.FileHeader
}

// Open implements Source.
func (i *innerMultipartSource) Open() (io.ReadCloser, error) {

	file, err := i.header.Open()

	if err != nil {
		return nil, err
	}

	return file, nil
}

////////////////////////////////////////////////////////////////////////////////
