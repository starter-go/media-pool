package client

import (
	"github.com/starter-go/application"
	"github.com/starter-go/media-pool/common/classes/objects"
)

// the media-pool-client
type MPClient struct {

	//starter:component

	_as func(objects.Client) //starter:as("#")

	FilterList []objects.FilterRegistry //starter:inject(".")

	chainUp   objects.UploadFilterChain
	chainDown objects.DownloadFilterChain
}

// Fetch implements objects.Client.
func (inst *MPClient) Fetch(o *objects.Object) error {

	return inst.chainDown.Fetch(o)
}

// Put implements objects.Client.
func (inst *MPClient) Put(o *objects.Object) error {

	return inst.chainUp.Put(o)
}

func (inst *MPClient) Life() *application.Life {
	l := new(application.Life)
	l.OnCreate = inst.onCreate
	return l
}

func (inst *MPClient) onCreate() error {
	return inst.innerLoadChain()
}

func (inst *MPClient) innerLoadChain() error {

	ldr := new(objects.FilterChainLoader)
	ldr.Init(inst.FilterList)

	inst.chainUp = ldr.LoadChainUp()
	inst.chainDown = ldr.LoadChainDown()
	return nil
}

func (inst *MPClient) _impl() objects.Client {
	return inst
}
