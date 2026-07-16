package server4mediapool
import (
    p0d2a11d16 "github.com/starter-go/afs"
    p262c04a06 "github.com/starter-go/buckets"
    pd1a916a20 "github.com/starter-go/libgin"
    p0bd1dba0d "github.com/starter-go/media-pool/common/classes/caches"
    p46220f67d "github.com/starter-go/media-pool/common/classes/objects"
    p4868bf213 "github.com/starter-go/media-pool/common/classes/pools"
    pdfc4d7922 "github.com/starter-go/media-pool/server"
    p1308d9a33 "github.com/starter-go/media-pool/server/implements/icache"
    p821808f49 "github.com/starter-go/media-pool/server/implements/ifiles"
    pa1ce7a44d "github.com/starter-go/media-pool/server/implements/ihash"
    p7aa71734e "github.com/starter-go/media-pool/server/implements/imeta"
    pa7b549037 "github.com/starter-go/media-pool/server/implements/imonitor"
    p8d6dba9bf "github.com/starter-go/media-pool/server/implements/iobjects"
    p684d1b16e "github.com/starter-go/media-pool/server/implements/ipath"
    pfc90c21d1 "github.com/starter-go/media-pool/server/implements/istorage"
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



// type p1308d9a33.CacheDirAgentImpl in package:github.com/starter-go/media-pool/server/implements/icache
//
// id:com-1308d9a3324de391-icache-CacheDirAgentImpl
// class:
// alias:alias-0bd1dba0d0dee7ca0d5666cdcbd63313-CacheDirAgent
// scope:singleton
//
type p1308d9a332_icache_CacheDirAgentImpl struct {
}

func (inst* p1308d9a332_icache_CacheDirAgentImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1308d9a3324de391-icache-CacheDirAgentImpl"
	r.Classes = ""
	r.Aliases = "alias-0bd1dba0d0dee7ca0d5666cdcbd63313-CacheDirAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1308d9a332_icache_CacheDirAgentImpl) new() any {
    return &p1308d9a33.CacheDirAgentImpl{}
}

func (inst* p1308d9a332_icache_CacheDirAgentImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1308d9a33.CacheDirAgentImpl)
	nop(ie, com)

	
    com.MPCacheDir = inst.getMPCacheDir(ie)
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p1308d9a332_icache_CacheDirAgentImpl) getMPCacheDir(ie application.InjectionExt)string{
    return ie.GetString("${mediapool.cache.dir}")
}


func (inst*p1308d9a332_icache_CacheDirAgentImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}



// type p1308d9a33.ObjectCacheFilterLayer in package:github.com/starter-go/media-pool/server/implements/icache
//
// id:com-1308d9a3324de391-icache-ObjectCacheFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p1308d9a332_icache_ObjectCacheFilterLayer struct {
}

func (inst* p1308d9a332_icache_ObjectCacheFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1308d9a3324de391-icache-ObjectCacheFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1308d9a332_icache_ObjectCacheFilterLayer) new() any {
    return &p1308d9a33.ObjectCacheFilterLayer{}
}

func (inst* p1308d9a332_icache_ObjectCacheFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1308d9a33.ObjectCacheFilterLayer)
	nop(ie, com)

	


    return nil
}



// type p821808f49.FileSetFilter in package:github.com/starter-go/media-pool/server/implements/ifiles
//
// id:com-821808f49626e9e9-ifiles-FileSetFilter
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p821808f496_ifiles_FileSetFilter struct {
}

func (inst* p821808f496_ifiles_FileSetFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-821808f49626e9e9-ifiles-FileSetFilter"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p821808f496_ifiles_FileSetFilter) new() any {
    return &p821808f49.FileSetFilter{}
}

func (inst* p821808f496_ifiles_FileSetFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p821808f49.FileSetFilter)
	nop(ie, com)

	
    com.CDAgent = inst.getCDAgent(ie)


    return nil
}


func (inst*p821808f496_ifiles_FileSetFilter) getCDAgent(ie application.InjectionExt)p0bd1dba0d.CacheDirAgent{
    return ie.GetComponent("#alias-0bd1dba0d0dee7ca0d5666cdcbd63313-CacheDirAgent").(p0bd1dba0d.CacheDirAgent)
}



// type p821808f49.TempFileFilterLayer in package:github.com/starter-go/media-pool/server/implements/ifiles
//
// id:com-821808f49626e9e9-ifiles-TempFileFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p821808f496_ifiles_TempFileFilterLayer struct {
}

func (inst* p821808f496_ifiles_TempFileFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-821808f49626e9e9-ifiles-TempFileFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p821808f496_ifiles_TempFileFilterLayer) new() any {
    return &p821808f49.TempFileFilterLayer{}
}

func (inst* p821808f496_ifiles_TempFileFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p821808f49.TempFileFilterLayer)
	nop(ie, com)

	
    com.CDAgent = inst.getCDAgent(ie)


    return nil
}


func (inst*p821808f496_ifiles_TempFileFilterLayer) getCDAgent(ie application.InjectionExt)p0bd1dba0d.CacheDirAgent{
    return ie.GetComponent("#alias-0bd1dba0d0dee7ca0d5666cdcbd63313-CacheDirAgent").(p0bd1dba0d.CacheDirAgent)
}



// type pa1ce7a44d.SumFilterLayer in package:github.com/starter-go/media-pool/server/implements/ihash
//
// id:com-a1ce7a44dba3bf2e-ihash-SumFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type pa1ce7a44db_ihash_SumFilterLayer struct {
}

func (inst* pa1ce7a44db_ihash_SumFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a1ce7a44dba3bf2e-ihash-SumFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa1ce7a44db_ihash_SumFilterLayer) new() any {
    return &pa1ce7a44d.SumFilterLayer{}
}

func (inst* pa1ce7a44db_ihash_SumFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa1ce7a44d.SumFilterLayer)
	nop(ie, com)

	


    return nil
}



// type p7aa71734e.MetaFilter in package:github.com/starter-go/media-pool/server/implements/imeta
//
// id:com-7aa71734e5dc3492-imeta-MetaFilter
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p7aa71734e5_imeta_MetaFilter struct {
}

func (inst* p7aa71734e5_imeta_MetaFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-7aa71734e5dc3492-imeta-MetaFilter"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7aa71734e5_imeta_MetaFilter) new() any {
    return &p7aa71734e.MetaFilter{}
}

func (inst* p7aa71734e5_imeta_MetaFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p7aa71734e.MetaFilter)
	nop(ie, com)

	


    return nil
}



// type pa7b549037.MonitorFilterLayer in package:github.com/starter-go/media-pool/server/implements/imonitor
//
// id:com-a7b549037c5641a8-imonitor-MonitorFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type pa7b549037c_imonitor_MonitorFilterLayer struct {
}

func (inst* pa7b549037c_imonitor_MonitorFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a7b549037c5641a8-imonitor-MonitorFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa7b549037c_imonitor_MonitorFilterLayer) new() any {
    return &pa7b549037.MonitorFilterLayer{}
}

func (inst* pa7b549037c_imonitor_MonitorFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa7b549037.MonitorFilterLayer)
	nop(ie, com)

	


    return nil
}



// type pa7b549037.ParamsCheckerFilter in package:github.com/starter-go/media-pool/server/implements/imonitor
//
// id:com-a7b549037c5641a8-imonitor-ParamsCheckerFilter
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type pa7b549037c_imonitor_ParamsCheckerFilter struct {
}

func (inst* pa7b549037c_imonitor_ParamsCheckerFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a7b549037c5641a8-imonitor-ParamsCheckerFilter"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa7b549037c_imonitor_ParamsCheckerFilter) new() any {
    return &pa7b549037.ParamsCheckerFilter{}
}

func (inst* pa7b549037c_imonitor_ParamsCheckerFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa7b549037.ParamsCheckerFilter)
	nop(ie, com)

	
    com.MaxContentLength = inst.getMaxContentLength(ie)
    com.MinContentLength = inst.getMinContentLength(ie)


    return nil
}


func (inst*pa7b549037c_imonitor_ParamsCheckerFilter) getMaxContentLength(ie application.InjectionExt)int64{
    return ie.GetInt64("${mediapool.web.max-content-length}")
}


func (inst*pa7b549037c_imonitor_ParamsCheckerFilter) getMinContentLength(ie application.InjectionExt)int64{
    return ie.GetInt64("${mediapool.web.min-content-length}")
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



// type p684d1b16e.PathMakerFilterLayer in package:github.com/starter-go/media-pool/server/implements/ipath
//
// id:com-684d1b16e2ad1f5e-ipath-PathMakerFilterLayer
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p684d1b16e2_ipath_PathMakerFilterLayer struct {
}

func (inst* p684d1b16e2_ipath_PathMakerFilterLayer) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-684d1b16e2ad1f5e-ipath-PathMakerFilterLayer"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p684d1b16e2_ipath_PathMakerFilterLayer) new() any {
    return &p684d1b16e.PathMakerFilterLayer{}
}

func (inst* p684d1b16e2_ipath_PathMakerFilterLayer) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p684d1b16e.PathMakerFilterLayer)
	nop(ie, com)

	


    return nil
}



// type p684d1b16e.WebLocationMakerFilter in package:github.com/starter-go/media-pool/server/implements/ipath
//
// id:com-684d1b16e2ad1f5e-ipath-WebLocationMakerFilter
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type p684d1b16e2_ipath_WebLocationMakerFilter struct {
}

func (inst* p684d1b16e2_ipath_WebLocationMakerFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-684d1b16e2ad1f5e-ipath-WebLocationMakerFilter"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p684d1b16e2_ipath_WebLocationMakerFilter) new() any {
    return &p684d1b16e.WebLocationMakerFilter{}
}

func (inst* p684d1b16e2_ipath_WebLocationMakerFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p684d1b16e.WebLocationMakerFilter)
	nop(ie, com)

	
    com.RawBaseObjectURL = inst.getRawBaseObjectURL(ie)


    return nil
}


func (inst*p684d1b16e2_ipath_WebLocationMakerFilter) getRawBaseObjectURL(ie application.InjectionExt)string{
    return ie.GetString("${mediapool.web.objects-base-url}")
}



// type pfc90c21d1.DefaultBucketHolder in package:github.com/starter-go/media-pool/server/implements/istorage
//
// id:com-fc90c21d18000bd9-istorage-DefaultBucketHolder
// class:
// alias:alias-4868bf21398de71b737baf5a174a3530-BucketHolder
// scope:singleton
//
type pfc90c21d18_istorage_DefaultBucketHolder struct {
}

func (inst* pfc90c21d18_istorage_DefaultBucketHolder) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fc90c21d18000bd9-istorage-DefaultBucketHolder"
	r.Classes = ""
	r.Aliases = "alias-4868bf21398de71b737baf5a174a3530-BucketHolder"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfc90c21d18_istorage_DefaultBucketHolder) new() any {
    return &pfc90c21d1.DefaultBucketHolder{}
}

func (inst* pfc90c21d18_istorage_DefaultBucketHolder) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfc90c21d1.DefaultBucketHolder)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*pfc90c21d18_istorage_DefaultBucketHolder) getLoader(ie application.InjectionExt)p4868bf213.BucketLoader{
    return ie.GetComponent("#alias-4868bf21398de71b737baf5a174a3530-BucketLoader").(p4868bf213.BucketLoader)
}



// type pfc90c21d1.DefaultBucketLoader in package:github.com/starter-go/media-pool/server/implements/istorage
//
// id:com-fc90c21d18000bd9-istorage-DefaultBucketLoader
// class:
// alias:alias-4868bf21398de71b737baf5a174a3530-BucketLoader
// scope:singleton
//
type pfc90c21d18_istorage_DefaultBucketLoader struct {
}

func (inst* pfc90c21d18_istorage_DefaultBucketLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fc90c21d18000bd9-istorage-DefaultBucketLoader"
	r.Classes = ""
	r.Aliases = "alias-4868bf21398de71b737baf5a174a3530-BucketLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfc90c21d18_istorage_DefaultBucketLoader) new() any {
    return &pfc90c21d1.DefaultBucketLoader{}
}

func (inst* pfc90c21d18_istorage_DefaultBucketLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfc90c21d1.DefaultBucketLoader)
	nop(ie, com)

	
    com.BucketSer = inst.getBucketSer(ie)
    com.BucketName = inst.getBucketName(ie)


    return nil
}


func (inst*pfc90c21d18_istorage_DefaultBucketLoader) getBucketSer(ie application.InjectionExt)p262c04a06.Service{
    return ie.GetComponent("#alias-262c04a06c32904104382e2b8d56c279-Service").(p262c04a06.Service)
}


func (inst*pfc90c21d18_istorage_DefaultBucketLoader) getBucketName(ie application.InjectionExt)string{
    return ie.GetString("${mediapool.bucket.name}")
}



// type pfc90c21d1.ObjectStoragePoolFilter in package:github.com/starter-go/media-pool/server/implements/istorage
//
// id:com-fc90c21d18000bd9-istorage-ObjectStoragePoolFilter
// class:class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry
// alias:
// scope:singleton
//
type pfc90c21d18_istorage_ObjectStoragePoolFilter struct {
}

func (inst* pfc90c21d18_istorage_ObjectStoragePoolFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fc90c21d18000bd9-istorage-ObjectStoragePoolFilter"
	r.Classes = "class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfc90c21d18_istorage_ObjectStoragePoolFilter) new() any {
    return &pfc90c21d1.ObjectStoragePoolFilter{}
}

func (inst* pfc90c21d18_istorage_ObjectStoragePoolFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfc90c21d1.ObjectStoragePoolFilter)
	nop(ie, com)

	
    com.BH = inst.getBH(ie)


    return nil
}


func (inst*pfc90c21d18_istorage_ObjectStoragePoolFilter) getBH(ie application.InjectionExt)p4868bf213.BucketHolder{
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



// type p474aadc6b.ObjectDownloadController in package:github.com/starter-go/media-pool/server/web/controllers
//
// id:com-474aadc6bee2144f-controllers-ObjectDownloadController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p474aadc6be_controllers_ObjectDownloadController struct {
}

func (inst* p474aadc6be_controllers_ObjectDownloadController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-474aadc6bee2144f-controllers-ObjectDownloadController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p474aadc6be_controllers_ObjectDownloadController) new() any {
    return &p474aadc6b.ObjectDownloadController{}
}

func (inst* p474aadc6be_controllers_ObjectDownloadController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p474aadc6b.ObjectDownloadController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p474aadc6be_controllers_ObjectDownloadController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p474aadc6be_controllers_ObjectDownloadController) getService(ie application.InjectionExt)p46220f67d.Server{
    return ie.GetComponent("#alias-46220f67d06e6dbd28c3603d4b14f6ae-Server").(p46220f67d.Server)
}



// type p474aadc6b.ObjectUploadController in package:github.com/starter-go/media-pool/server/web/controllers
//
// id:com-474aadc6bee2144f-controllers-ObjectUploadController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p474aadc6be_controllers_ObjectUploadController struct {
}

func (inst* p474aadc6be_controllers_ObjectUploadController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-474aadc6bee2144f-controllers-ObjectUploadController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p474aadc6be_controllers_ObjectUploadController) new() any {
    return &p474aadc6b.ObjectUploadController{}
}

func (inst* p474aadc6be_controllers_ObjectUploadController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p474aadc6b.ObjectUploadController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p474aadc6be_controllers_ObjectUploadController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p474aadc6be_controllers_ObjectUploadController) getService(ie application.InjectionExt)p46220f67d.Server{
    return ie.GetComponent("#alias-46220f67d06e6dbd28c3603d4b14f6ae-Server").(p46220f67d.Server)
}


