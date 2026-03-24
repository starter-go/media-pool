package objects

////////////////////////////////////////////////////////////////////////////////

type UploadFilter interface {
	Put(o *Object, next UploadFilterChain) error
}

type UploadFilterChain interface {
	Put(o *Object) error
}

////////////////////////////////////////////////////////////////////////////////

type DownloadFilter interface {
	Fetch(o *Object, next DownloadFilterChain) error
}

type DownloadFilterChain interface {
	Fetch(o *Object) error
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
