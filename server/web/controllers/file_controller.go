package controllers

import (
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/classes/streams"
	"github.com/starter-go/media-pool/common/data/dxo"
	"github.com/starter-go/media-pool/server/web/dto"
	"github.com/starter-go/media-pool/server/web/vo"

	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type MediaController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")

	Service objects.Server //starter:inject("#")

}

func (inst *MediaController) _impl() libgin.Controller {
	return inst
}

func (inst *MediaController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *MediaController) route(rp libgin.RouterProxy) error {

	rp = rp.For("/files")
	path := ":type1/:type2/:id/:name"

	//fetch

	rp.GET(path, inst.handleGetOne)

	// put

	rp.POST(path, inst.handlePostUpload)
	rp.POST("upload", inst.handlePostUpload)

	return nil
}

func (inst *MediaController) handleGetOne(gc *gin.Context) {

	req := new(myMediaFetchRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetOne)
}

func (inst *MediaController) handlePostUpload(gc *gin.Context) {

	req := new(myMediaUploadRequest)
	req.context = gc
	req.controller = inst

	req.execute(req.handleUploadFile)
}

func (inst *MediaController) handleGetList(gc *gin.Context) {

	req := new(myMediaFetchRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = false
	req.wantRequestBody = false

	req.execute(req.doGetList)
}

func (inst *MediaController) handlePutItem(gc *gin.Context) {

	req := new(myMediaFetchRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestID = true
	req.wantRequestBody = true

	req.execute(req.doPutItem)
}

////////////////////////////////////////////////////////////////////////////////

type myMediaFetchRequest struct {
	wantRequestID   bool
	wantRequestBody bool

	context    *gin.Context
	controller *MediaController

	id    dxo.ExampleID
	body1 vo.Examples
	body2 vo.Examples
}

func (inst *myMediaFetchRequest) open(ctx *gin.Context) error {

	if inst.wantRequestID {
		str := ctx.Param("id")
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		inst.id = dxo.ExampleID(num)
	}

	if inst.wantRequestBody {
		obj := &inst.body1
		err := ctx.BindJSON(obj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myMediaFetchRequest) execute(task func() error) {

	ex := new(libgin.Executor)
	ex.Context = inst.context
	ex.Responder = inst.controller.Sender
	ex.Body1 = &inst.body1
	ex.Body2 = &inst.body2

	ex.OnOpen = inst.open
	ex.OnTask = task

	ex.Execute()
}

func (inst *myMediaFetchRequest) doGetList() error {

	it := &dto.Example{}

	inst.body2.Items = []*dto.Example{it, it, it}
	return nil
}

func (inst *myMediaFetchRequest) doGetOne() error {

	ser := inst.controller.Service

	o := &objects.Info{
		ID:   "",
		Name: "",
		Type: "",
		Data: nil,
	}

	ser.Fetch(o)

	return nil
}

func (inst *myMediaFetchRequest) doPutItem() error {

	it1 := inst.body1.Items[0]
	it2 := &dto.Example{}
	id := inst.id

	it2.ID = id

	inst.body2.Items = []*dto.Example{it1, it2}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// type  myMediaUploadRequest

type myMediaUploadRequest struct {
	context    *gin.Context
	controller *MediaController

	items []*myMediaUploadFile
}

func (inst *myMediaUploadRequest) execute(fn func() error) {

	//  read request
	err := inst.open()

	// handle data
	if err == nil {
		err = fn()
	}

	// send result
	inst.send(err)
}

func (inst *myMediaUploadRequest) open() error {

	c := inst.context
	mpf, err := c.MultipartForm()
	if err != nil {
		return err
	}

	table1 := mpf.File

	for k1, v1 := range table1 {

		list2 := v1

		for _, item2 := range list2 {

			// size := item2.Size
			// name := item2.Filename
			// vlog.Debug(" handle uploading file:   key=%s ; filename=%s ; size=%d", k1, name, size)

			item3 := new(myMediaUploadFile)
			item3.key = k1
			item3.init(item2)
			inst.items = append(inst.items, item3)
		}
	}

	return nil
}

func (inst *myMediaUploadRequest) send(err error) {

	ctx := inst.context
	sender := inst.controller.Sender

	resp := &libgin.Response{
		Context: ctx,
		Error:   err,
	}

	sender.Send(resp)

}

func (inst *myMediaUploadRequest) handleUploadFile() error {

	ser := inst.controller.Service
	items := inst.items
	ctx := inst.context

	for _, item := range items {

		d := item.getData()

		o := &objects.Object{
			Context: ctx,
			Size:    item.size,
			Name:    item.filename,
			Type:    item.contentType,
			Data:    d,
		}

		err := ser.Put(o)
		if err != nil {
			return err
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// myMediaUploadFile

type myMediaUploadFile struct {
	key         string
	filename    string
	contentType string
	size        int64

	header *multipart.FileHeader
}

func (inst *myMediaUploadFile) init(h *multipart.FileHeader) {

	inst.size = h.Size
	inst.filename = h.Filename
	inst.header = h
	inst.contentType = h.Header.Get("content-type")

	key := inst.key
	name := inst.filename
	size := inst.size
	tp := inst.contentType

	vlog.Debug(" handle uploading file:   key=%s; filename=%s; size=%d; type=%s", key, name, size, tp)

}

func (inst *myMediaUploadFile) getData() streams.Source {
	return streams.NewSourceForMultipart(inst.header)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
