package buffers

import (
	"io"

	"github.com/starter-go/afs"
	"github.com/starter-go/media-pool/common/classes/streams"
)

type Buffer interface {
	io.Closer

	streams.Source

	GetSize() int

	OpenWriter() (io.WriteCloser, error)

	OpenReader() (io.ReadCloser, error)

	// 如果是基于文件实现的缓冲区, 返回文件的位置
	GetPath() afs.Path
}
