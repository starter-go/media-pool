package server4mediapool

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p474aadc6be_controllers_ExampleController{})
    inst.register(&p474aadc6be_controllers_MediaController{})
    inst.register(&p8d6dba9bfd_iobjects_LocalObjectCache{})
    inst.register(&p8d6dba9bfd_iobjects_ObjectServiceImpl{})
    inst.register(&pdfc4d7922e_server_MPServer{})
    inst.register(&pe2332cb1bd_ipools_DefaultBucketHolder{})
    inst.register(&pe2332cb1bd_ipools_DefaultBucketLoader{})
    inst.register(&pe2332cb1bd_ipools_DefaultPoolService{})


    return nil
}
