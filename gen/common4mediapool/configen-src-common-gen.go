package common4mediapool
import (
    p0d2a11d16 "github.com/starter-go/afs"
    p78c4450e8 "github.com/starter-go/media-pool/common/impl/icache"
    pc2ffe7639 "github.com/starter-go/media-pool/common/impl/ihash"
    p33805c9ff "github.com/starter-go/media-pool/common/impl/imonitor"
    peab988321 "github.com/starter-go/media-pool/common/impl/ipath"
    p72ff7347b "github.com/starter-go/media-pool/common/impl/itempfile"
     "github.com/starter-go/application"
)

// type p78c4450e8.ObjectCacheFilterLayer in package:github.com/starter-go/media-pool/common/impl/icache
//
// id:com-78c4450e8d44ee02-icache-ObjectCacheFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p78c4450e8d_icache_ObjectCacheFilterLayer struct {
}

func (inst* p78c4450e8d_icache_ObjectCacheFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-78c4450e8d44ee02-icache-ObjectCacheFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p78c4450e8d_icache_ObjectCacheFilterLayer) new() any {
    return &p78c4450e8.ObjectCacheFilterLayer{}
}

func (inst* p78c4450e8d_icache_ObjectCacheFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p78c4450e8.ObjectCacheFilterLayer)
	nop(ie, com)

	
    com.MPCacheDir = inst.getMPCacheDir(ie)
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p78c4450e8d_icache_ObjectCacheFilterLayer) getMPCacheDir(ie application.InjectionExt)string{
    return ie.GetString("${mediapool.cache.dir}")
}


func (inst*p78c4450e8d_icache_ObjectCacheFilterLayer) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type pc2ffe7639.SumFilterLayer in package:github.com/starter-go/media-pool/common/impl/ihash
//
// id:com-c2ffe763909db52f-ihash-SumFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type pc2ffe76390_ihash_SumFilterLayer struct {
}

func (inst* pc2ffe76390_ihash_SumFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c2ffe763909db52f-ihash-SumFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc2ffe76390_ihash_SumFilterLayer) new() any {
    return &pc2ffe7639.SumFilterLayer{}
}

func (inst* pc2ffe76390_ihash_SumFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc2ffe7639.SumFilterLayer)
	nop(ie, com)

	


    return nil
}



// type p33805c9ff.MonitorFilterLayer in package:github.com/starter-go/media-pool/common/impl/imonitor
//
// id:com-33805c9ff5ec1ca4-imonitor-MonitorFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p33805c9ff5_imonitor_MonitorFilterLayer struct {
}

func (inst* p33805c9ff5_imonitor_MonitorFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-33805c9ff5ec1ca4-imonitor-MonitorFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p33805c9ff5_imonitor_MonitorFilterLayer) new() any {
    return &p33805c9ff.MonitorFilterLayer{}
}

func (inst* p33805c9ff5_imonitor_MonitorFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p33805c9ff.MonitorFilterLayer)
	nop(ie, com)

	


    return nil
}



// type p33805c9ff.ParamsCheckerFilter in package:github.com/starter-go/media-pool/common/impl/imonitor
//
// id:com-33805c9ff5ec1ca4-imonitor-ParamsCheckerFilter
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p33805c9ff5_imonitor_ParamsCheckerFilter struct {
}

func (inst* p33805c9ff5_imonitor_ParamsCheckerFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-33805c9ff5ec1ca4-imonitor-ParamsCheckerFilter"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p33805c9ff5_imonitor_ParamsCheckerFilter) new() any {
    return &p33805c9ff.ParamsCheckerFilter{}
}

func (inst* p33805c9ff5_imonitor_ParamsCheckerFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p33805c9ff.ParamsCheckerFilter)
	nop(ie, com)

	


    return nil
}



// type peab988321.PathMakerFilterLayer in package:github.com/starter-go/media-pool/common/impl/ipath
//
// id:com-eab9883210a51acc-ipath-PathMakerFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type peab9883210_ipath_PathMakerFilterLayer struct {
}

func (inst* peab9883210_ipath_PathMakerFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eab9883210a51acc-ipath-PathMakerFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peab9883210_ipath_PathMakerFilterLayer) new() any {
    return &peab988321.PathMakerFilterLayer{}
}

func (inst* peab9883210_ipath_PathMakerFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peab988321.PathMakerFilterLayer)
	nop(ie, com)

	


    return nil
}



// type p72ff7347b.TempFileFilterLayer in package:github.com/starter-go/media-pool/common/impl/itempfile
//
// id:com-72ff7347bb4ab226-itempfile-TempFileFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p72ff7347bb_itempfile_TempFileFilterLayer struct {
}

func (inst* p72ff7347bb_itempfile_TempFileFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72ff7347bb4ab226-itempfile-TempFileFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p72ff7347bb_itempfile_TempFileFilterLayer) new() any {
    return &p72ff7347b.TempFileFilterLayer{}
}

func (inst* p72ff7347bb_itempfile_TempFileFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p72ff7347b.TempFileFilterLayer)
	nop(ie, com)

	


    return nil
}


