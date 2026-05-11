package objects

////////////////////////////////////////////////////////////////////////////////

type UploadFilter interface {
	Put(o *IOContext, next UploadFilterChain) error
}

type UploadFilterChain interface {
	Put(o *IOContext) error
}

////////////////////////////////////////////////////////////////////////////////

type DownloadFilter interface {
	Fetch(o *IOContext, next DownloadFilterChain) error
}

type DownloadFilterChain interface {
	Fetch(o *IOContext) error
}

////////////////////////////////////////////////////////////////////////////////

type FilterRegistration struct {
	Enabled bool

	Priority int

	Label string

	Class string // a list of classes, like "aa bb cc"

	Down DownloadFilter

	Up UploadFilter
}

type FilterRegistry interface {
	ListFilters() []*FilterRegistration
}

////////////////////////////////////////////////////////////////////////////////
// EOF
