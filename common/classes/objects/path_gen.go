package objects

type PathGenerator interface {
	GenPath(o *Info) Path
}
