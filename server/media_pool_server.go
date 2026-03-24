package server

import (
	"github.com/starter-go/application"
	"github.com/starter-go/media-pool/common/classes/objects"
)

// the media-pool-server
type MPServer struct {

	//starter:component

	_as func(objects.Server) //starter:as("#")

	FilterList []objects.FilterRegistry //starter:inject(".")

	chainUp   objects.UploadFilterChain
	chainDown objects.DownloadFilterChain
}

// Fetch implements objects.Client.
func (inst *MPServer) Fetch(o *objects.Object) error {

	return inst.chainDown.Fetch(o)
}

// Put implements objects.Client.
func (inst *MPServer) Put(o *objects.Object) error {

	return inst.chainUp.Put(o)
}

func (inst *MPServer) Life() *application.Life {
	l := new(application.Life)
	l.OnCreate = inst.onCreate
	return l
}

func (inst *MPServer) onCreate() error {
	return inst.innerLoadChain()
}

func (inst *MPServer) innerLoadChain() error {

	ldr := new(objects.FilterChainLoader)
	ldr.Init(inst.FilterList)

	inst.chainUp = ldr.LoadChainUp()
	inst.chainDown = ldr.LoadChainDown()
	return nil
}

func (inst *MPServer) _impl() objects.Client {
	return inst
}
