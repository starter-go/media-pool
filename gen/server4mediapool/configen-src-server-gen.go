package server4mediapool
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p46220f67d "github.com/starter-go/media-pool/common/classes/objects"
    p4868bf213 "github.com/starter-go/media-pool/common/classes/pools"
    pdfc4d7922 "github.com/starter-go/media-pool/server"
    p8d6dba9bf "github.com/starter-go/media-pool/server/implements/iobjects"
    pe2332cb1b "github.com/starter-go/media-pool/server/implements/ipools"
    p474aadc6b "github.com/starter-go/media-pool/server/web/controllers"
     "github.com/starter-go/application"
)

// type pdfc4d7922.MPServer in package:github.com/starter-go/media-pool/server
//
// id:com-dfc4d7922ec88149-server-MPServer
// class:
// alias:alias-46220f67d06e6dbd28c3603d4b14f6ae-Server
// scope:singleton
//
type pdfc4d7922e_server_MPServer struct {
}

func (inst* pdfc4d7922e_server_MPServer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-dfc4d7922ec88149-server-MPServer"
	r.Classes = ""
	r.Aliases = "alias-46220f67d06e6dbd28c3603d4b14f6ae-Server"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdfc4d7922e_server_MPServer) new() any {
    return &pdfc4d7922.MPServer{}
}

func (inst* pdfc4d7922e_server_MPServer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdfc4d7922.MPServer)
	nop(ie, com)

	
    com.FilterList = inst.getFilterList(ie)


    return nil
}


func (inst*pdfc4d7922e_server_MPServer) getFilterList(ie application.InjectionExt)[]p46220f67d.FilterRegistry{
    dst := make([]p46220f67d.FilterRegistry, 0)
    src := ie.ListComponents(".class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p46220f67d.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p8d6dba9bf.LocalObjectCache in package:github.com/starter-go/media-pool/server/implements/iobjects
//
// id:com-8d6dba9bfdabfebf-iobjects-LocalObjectCache
// class:
// alias:alias-0bd1dba0d0dee7ca0d5666cdcbd63313-Service
// scope:singleton
//
type p8d6dba9bfd_iobjects_LocalObjectCache struct {
}

func (inst* p8d6dba9bfd_iobjects_LocalObjectCache) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d6dba9bfdabfebf-iobjects-LocalObjectCache"
	r.Classes = ""
	r.Aliases = "alias-0bd1dba0d0dee7ca0d5666cdcbd63313-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d6dba9bfd_iobjects_LocalObjectCache) new() any {
    return &p8d6dba9bf.LocalObjectCache{}
}

func (inst* p8d6dba9bfd_iobjects_LocalObjectCache) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d6dba9bf.LocalObjectCache)
	nop(ie, com)

	


    return nil
}



// type p8d6dba9bf.ObjectServiceImpl in package:github.com/starter-go/media-pool/server/implements/iobjects
//
// id:com-8d6dba9bfdabfebf-iobjects-ObjectServiceImpl
// class:
// alias:alias-46220f67d06e6dbd28c3603d4b14f6ae-Service
// scope:singleton
//
type p8d6dba9bfd_iobjects_ObjectServiceImpl struct {
}

func (inst* p8d6dba9bfd_iobjects_ObjectServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8d6dba9bfdabfebf-iobjects-ObjectServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-46220f67d06e6dbd28c3603d4b14f6ae-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8d6dba9bfd_iobjects_ObjectServiceImpl) new() any {
    return &p8d6dba9bf.ObjectServiceImpl{}
}

func (inst* p8d6dba9bfd_iobjects_ObjectServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8d6dba9bf.ObjectServiceImpl)
	nop(ie, com)

	


    return nil
}



// type pe2332cb1b.DefaultBucketHolder in package:github.com/starter-go/media-pool/server/implements/ipools
//
// id:com-e2332cb1bdc14a93-ipools-DefaultBucketHolder
// class:
// alias:alias-4868bf21398de71b737baf5a174a3530-BucketHolder
// scope:singleton
//
type pe2332cb1bd_ipools_DefaultBucketHolder struct {
}

func (inst* pe2332cb1bd_ipools_DefaultBucketHolder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e2332cb1bdc14a93-ipools-DefaultBucketHolder"
	r.Classes = ""
	r.Aliases = "alias-4868bf21398de71b737baf5a174a3530-BucketHolder"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe2332cb1bd_ipools_DefaultBucketHolder) new() any {
    return &pe2332cb1b.DefaultBucketHolder{}
}

func (inst* pe2332cb1bd_ipools_DefaultBucketHolder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe2332cb1b.DefaultBucketHolder)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*pe2332cb1bd_ipools_DefaultBucketHolder) getLoader(ie application.InjectionExt)p4868bf213.BucketLoader{
    return ie.GetComponent("#alias-4868bf21398de71b737baf5a174a3530-BucketLoader").(p4868bf213.BucketLoader)
}



// type pe2332cb1b.DefaultBucketLoader in package:github.com/starter-go/media-pool/server/implements/ipools
//
// id:com-e2332cb1bdc14a93-ipools-DefaultBucketLoader
// class:
// alias:alias-4868bf21398de71b737baf5a174a3530-BucketLoader
// scope:singleton
//
type pe2332cb1bd_ipools_DefaultBucketLoader struct {
}

func (inst* pe2332cb1bd_ipools_DefaultBucketLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e2332cb1bdc14a93-ipools-DefaultBucketLoader"
	r.Classes = ""
	r.Aliases = "alias-4868bf21398de71b737baf5a174a3530-BucketLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe2332cb1bd_ipools_DefaultBucketLoader) new() any {
    return &pe2332cb1b.DefaultBucketLoader{}
}

func (inst* pe2332cb1bd_ipools_DefaultBucketLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe2332cb1b.DefaultBucketLoader)
	nop(ie, com)

	


    return nil
}



// type pe2332cb1b.DefaultPoolService in package:github.com/starter-go/media-pool/server/implements/ipools
//
// id:com-e2332cb1bdc14a93-ipools-DefaultPoolService
// class:
// alias:alias-4868bf21398de71b737baf5a174a3530-Service
// scope:singleton
//
type pe2332cb1bd_ipools_DefaultPoolService struct {
}

func (inst* pe2332cb1bd_ipools_DefaultPoolService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e2332cb1bdc14a93-ipools-DefaultPoolService"
	r.Classes = ""
	r.Aliases = "alias-4868bf21398de71b737baf5a174a3530-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe2332cb1bd_ipools_DefaultPoolService) new() any {
    return &pe2332cb1b.DefaultPoolService{}
}

func (inst* pe2332cb1bd_ipools_DefaultPoolService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe2332cb1b.DefaultPoolService)
	nop(ie, com)

	
    com.BH = inst.getBH(ie)


    return nil
}


func (inst*pe2332cb1bd_ipools_DefaultPoolService) getBH(ie application.InjectionExt)p4868bf213.BucketHolder{
    return ie.GetComponent("#alias-4868bf21398de71b737baf5a174a3530-BucketHolder").(p4868bf213.BucketHolder)
}



// type p474aadc6b.ExampleController in package:github.com/starter-go/media-pool/server/web/controllers
//
// id:com-474aadc6bee2144f-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p474aadc6be_controllers_ExampleController struct {
}

func (inst* p474aadc6be_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-474aadc6bee2144f-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p474aadc6be_controllers_ExampleController) new() any {
    return &p474aadc6b.ExampleController{}
}

func (inst* p474aadc6be_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p474aadc6b.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*p474aadc6be_controllers_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type p474aadc6b.MediaController in package:github.com/starter-go/media-pool/server/web/controllers
//
// id:com-474aadc6bee2144f-controllers-MediaController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p474aadc6be_controllers_MediaController struct {
}

func (inst* p474aadc6be_controllers_MediaController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-474aadc6bee2144f-controllers-MediaController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p474aadc6be_controllers_MediaController) new() any {
    return &p474aadc6b.MediaController{}
}

func (inst* p474aadc6be_controllers_MediaController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p474aadc6b.MediaController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p474aadc6be_controllers_MediaController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p474aadc6be_controllers_MediaController) getService(ie application.InjectionExt)p46220f67d.Server{
    return ie.GetComponent("#alias-46220f67d06e6dbd28c3603d4b14f6ae-Server").(p46220f67d.Server)
}


