package streams

import "io"

type LengthLimiter interface {
	Limit(in io.ReadCloser, maxLen int) io.ReadCloser
}
