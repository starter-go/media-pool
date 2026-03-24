package objects

type IDGenerator interface {
	GenID(o *Info) ID
}
