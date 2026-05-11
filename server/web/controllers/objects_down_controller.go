package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/server/web/dto"
	"github.com/starter-go/media-pool/server/web/vo"
)

////////////////////////////////////////////////////////////////////////////////

type ObjectDownloadController struct {

	//starter:component

	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")

	Service objects.Server //starter:inject("#")

}

func (inst *ObjectDownloadController) _impl() libgin.Controller {
	return inst
}

func (inst *ObjectDownloadController) Registration() *libgin.ControllerRegistration {
	cr1 := new(libgin.ControllerRegistration)
	cr1.Route = inst.route
	return cr1
}

func (inst *ObjectDownloadController) route(rp libgin.RouterProxy) error {

	rp = rp.For("/objects")

	const path1a = ":id/:name" // default: auto

	// const path2m = ":type1/:type2/:id/meta"
	// const path3d = ":type1/:type2/:id/data"
	// const path4d = ":type1/:type2/:id/data/:name"
	// const path5t = ":type1/:type2/:id/thumbnail/:size"

	//fetch

	rp.GET(path1a, inst.handleGetOneAuto)

	// rp.GET(path2m, inst.handleGetOneMeta)
	// rp.GET(path3d, inst.handleGetOneData)
	// rp.GET(path4d, inst.handleGetOneData)
	// rp.GET(path5t, inst.handleGetOneThumb)

	return nil
}

func (inst *ObjectDownloadController) handleGetOneAuto(gc *gin.Context) {

	req := new(myMediaFetchRequest)
	req.context = gc
	req.controller = inst

	req.wantRequestIDs = true
	req.wantRequestQuery = true
	req.wantRequestBody = false

	req.execute(req.doGetOneAuto)
}

// func (inst *ObjectDownloadController) handleGetOneThumb(gc *gin.Context) {

// 	req := new(myMediaFetchRequest)
// 	req.context = gc
// 	req.controller = inst

// 	req.wantRequestIDs = true
// 	req.wantRequestBody = false

// 	req.execute(req.doGetOneThumb)
// }

// func (inst *ObjectDownloadController) handleGetOneData(gc *gin.Context) {

// 	req := new(myMediaFetchRequest)
// 	req.context = gc
// 	req.controller = inst

// 	req.wantRequestIDs = true
// 	req.wantRequestBody = false

// 	req.execute(req.doGetOneData)
// }

// func (inst *ObjectDownloadController) handleGetOneMeta(gc *gin.Context) {

// 	req := new(myMediaFetchRequest)
// 	req.context = gc
// 	req.controller = inst

// 	req.wantRequestIDs = true
// 	req.wantRequestBody = false

// 	req.execute(req.doGetOneMeta)
// }

////////////////////////////////////////////////////////////////////////////////

type myMediaFetchRequest struct {
	wantRequestIDs   bool
	wantRequestBody  bool
	wantRequestQuery bool

	wantResultMeta bool

	context    *gin.Context
	controller *ObjectDownloadController
	service    objects.Server

	paramID        objects.ID
	paramName      string
	paramType      string
	paramView      string // view = [raw|thumb|meta(default)|data]
	paramThumbSize int

	body1     vo.Objects
	body2     vo.Objects
	resultObj *objects.Object
	sent      bool
}

func (inst *myMediaFetchRequest) open(ctx *gin.Context) error {

	if inst.wantRequestIDs {

		strID := ctx.Param("id")
		strNa := ctx.Param("name")
		strT1 := ctx.Param("type1")
		strT2 := ctx.Param("type2")

		inst.paramID = objects.ID(strID)
		inst.paramName = strNa
		inst.paramType = strT1 + "/" + strT2
	}

	if inst.wantRequestQuery {
		strView := ctx.Query("view")
		strSize := ctx.Query("size") // for thumb
		nSize, _ := strconv.Atoi(strSize)
		inst.paramView = strView
		inst.paramThumbSize = nSize
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

	ctx := inst.context
	err := inst.open(ctx)

	if err == nil {
		err = task()
	}

	inst.sendError(err)
}

func (inst *myMediaFetchRequest) sendFileData(cf *objects.CacheFile) error {

	if inst.sent {
		return nil
	}

	if cf == nil {
		return fmt.Errorf("sendFileData() : objects.CacheFile is nil")
	}

	c := inst.context
	path := cf.File
	if path == nil {
		return fmt.Errorf("sendFileData() : file.path is nil")
	}

	code := http.StatusOK
	size1 := cf.ContentLength
	ctype := cf.ContentType

	info := path.GetInfo()
	size2 := info.Length()

	if size1 != size2 {
		return fmt.Errorf("sendFileData() : bad object-size (want=%d, have=%d)", size1, size2)
	}

	dataSrc, err := path.GetIO().OpenReader(nil)
	if err != nil {
		return err
	}
	defer dataSrc.Close()

	c.DataFromReader(code, size1, ctype, dataSrc, nil)

	inst.sent = true
	return nil
}

func (inst *myMediaFetchRequest) sendError(err error) {
	inst.sendJson(err)
}

func (inst *myMediaFetchRequest) sendJson(err error) {

	if inst.sent {
		return
	}

	body := &inst.body2
	code := body.Status
	now := lang.Now()

	if code == 0 {
		code = http.StatusOK
	}

	if err != nil {
		code = http.StatusInternalServerError
		body.Message = http.StatusText(code)
		body.Error = err.Error()
		body.Items = nil
	}

	body.Status = code
	body.Time = now.Time()
	body.Timestamp = now

	ctx := inst.context
	ctx.JSON(code, body)
	inst.sent = true
}

func (inst *myMediaFetchRequest) doGetOneAuto() error {

	view := inst.paramView

	if view == "raw" || view == "data" {
		return inst.doGetOneData()
	}

	if view == "thumb" {
		return inst.doGetOneThumb()
	}

	// default
	return inst.doGetOneMeta()
}

func (inst *myMediaFetchRequest) doGetOneThumb() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.paramID
	name := inst.paramName
	ctype := inst.paramType
	thumbSize := inst.paramThumbSize

	want := &objects.Info{
		Context:   ctx,
		Profile:   objects.ProfileThumb,
		ID:        id,
		Name:      name,
		Type:      ctype,
		Data:      nil,
		UseThumb:  true,
		ThumbSize: thumbSize,
	}
	ioc := &objects.IOContext{
		CC:   ctx,
		Want: want,
	}

	err := ser.Fetch(ioc)
	if err != nil {
		return err
	}

	have := ioc.Have
	fileSet := have.Files
	thumbFile := fileSet.Thumbnail256

	return inst.sendFileData(thumbFile)
}

func (inst *myMediaFetchRequest) doGetOneData() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.paramID
	name := inst.paramName
	ctype := inst.paramType

	want := &objects.Info{
		Context: ctx,
		Profile: objects.ProfileData,
		UseMeta: true,
		UseData: true,
		ID:      id,
		Name:    name,
		Type:    ctype,
		Data:    nil,
	}
	ioc := &objects.IOContext{
		CC:   ctx,
		Want: want,
	}

	err := ser.Fetch(ioc)
	if err != nil {
		return err
	}

	have := ioc.Have
	fileSet := &have.Files
	cf1d := fileSet.Data
	// cf2m := fileSet.Meta

	cf1d.ContentType = have.Type
	cf1d.ContentLength = have.Size

	return inst.sendFileData(cf1d)
}

func (inst *myMediaFetchRequest) innerConvertObject(src *objects.Object, dst *dto.Object) error {

	if src == nil || dst == nil {
		return fmt.Errorf("")
	}

	sum := src.Sum
	strSum := sum.String()

	dst.ID = src.ID
	dst.Name = src.Name
	dst.Type = src.Type
	dst.Length = src.Size
	dst.Sum = lang.Hex(strSum)
	dst.Type = src.Type

	return nil
}

func (inst *myMediaFetchRequest) doGetOneMeta() error {

	ctx := inst.context
	ser := inst.controller.Service
	id := inst.paramID
	name := inst.paramName
	ctype := inst.paramType

	want := &objects.Info{
		Context: ctx,
		Profile: objects.ProfileMeta,
		ID:      id,
		Name:    name,
		Type:    ctype,
		Data:    nil,
		UseMeta: true,
	}
	ioc := &objects.IOContext{
		CC:   ctx,
		Want: want,
	}

	err := ser.Fetch(ioc)
	if err != nil {
		return err
	}

	have1 := ioc.Have
	have2 := new(dto.Object)
	err = inst.innerConvertObject(have1, have2)
	if err != nil {
		return err
	}

	inst.body2.Items = []*dto.Object{have2}
	inst.sendJson(nil)
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
