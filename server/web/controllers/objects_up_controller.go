package controllers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/classes/streams"
	"github.com/starter-go/media-pool/server/web/dto"
	"github.com/starter-go/media-pool/server/web/vo"

	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type ObjectUploadController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")

	Service objects.Server //starter:inject("#")

}

func (inst *ObjectUploadController) _impl() libgin.Controller {
	return inst
}

func (inst *ObjectUploadController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *ObjectUploadController) route(rp libgin.RouterProxy) error {

	rp = rp.For("/objects")

	rp.POST("", inst.handlePostOne)
	rp.POST("upload", inst.handlePostOne)
	rp.POST(":type1/:type2/:id/:name", inst.handlePostOne)

	return nil
}

func (inst *ObjectUploadController) handlePostOne(gc *gin.Context) {

	req := new(myMediaUploadRequest)
	req.context = gc
	req.controller = inst

	req.execute(req.handleUploadFile)
}

////////////////////////////////////////////////////////////////////////////////
// type  myMediaUploadRequest

type myMediaUploadRequest struct {
	context *gin.Context

	controller *ObjectUploadController

	items []*myMediaUploadFile

	body2 vo.Objects
}

func (inst *myMediaUploadRequest) execute(fn func() error) {

	//  read request
	err := inst.open()

	// handle data
	if err == nil {
		err = fn()
	}

	// send result
	inst.sendJson(err)
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

func (inst *myMediaUploadRequest) sendJson(err error) {

	ctx := inst.context

	if err != nil {
		codeErr := http.StatusInternalServerError
		ctx.AbortWithError(codeErr, err)
		return
	}

	codeOk := http.StatusOK
	body := &inst.body2
	ctx.JSON(codeOk, body)
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
		ioc := &objects.IOContext{
			CC:   ctx,
			Want: o,
		}
		err := ser.Put(ioc)
		if err != nil {
			return err
		}
		inst.innerHandleUploadResult(ioc.Have)
	}

	return nil
}

func (inst *myMediaUploadRequest) innerHandleUploadResult(res *objects.Object) error {

	src := res
	dst := new(dto.Object)

	err := objects.Convert2DTO(src, dst)
	if err != nil {
		return err
	}

	list := inst.body2.Items
	list = append(list, dst)
	inst.body2.Items = list
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
