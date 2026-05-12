package common4mediapool

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

    
    inst.register(&p33805c9ff5_imonitor_MonitorFilterLayer{})
    inst.register(&p33805c9ff5_imonitor_ParamsCheckerFilter{})
    inst.register(&p72ff7347bb_itempfile_TempFileFilterLayer{})
    inst.register(&p78c4450e8d_icache_ObjectCacheFilterLayer{})
    inst.register(&pc2ffe76390_ihash_SumFilterLayer{})
    inst.register(&peab9883210_ipath_PathMakerFilterLayer{})
    inst.register(&peab9883210_ipath_WebLocationMakerFilter{})


    return nil
}
