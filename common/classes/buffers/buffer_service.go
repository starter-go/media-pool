package buffers

type Service interface {
	OpenBuffer(size int) (Buffer, error)
}
