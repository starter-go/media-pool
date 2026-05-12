package mediapool

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/bucket-drivers/aliyun"
	"github.com/starter-go/buckets/modules/buckets"
	"github.com/starter-go/libgin/modules/libgin"

	"github.com/starter-go/media-pool/gen/client4mediapool"
	"github.com/starter-go/media-pool/gen/common4mediapool"
	"github.com/starter-go/media-pool/gen/server4mediapool"
	"github.com/starter-go/media-pool/gen/test4mediapool"
	"github.com/starter-go/starter"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/starter-go/media-pool"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath   = "src/main/resources"
	theTestModuleResPath   = "src/test/resources"
	theClientModuleResPath = "src/client/resources"
	theServerModuleResPath = "src/server/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

//go:embed "src/client/resources"
var theClientModuleResFS embed.FS

//go:embed "src/server/resources"
var theServerModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func ModuleForCommon() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#common")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

	mb.Components(common4mediapool.ExportComponents)

	mb.Depend(starter.Module())

	return mb.Create()
}

func ModuleForClient() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#client")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theClientModuleResFS, theClientModuleResPath)

	mb.Components(client4mediapool.ExportComponents)

	mb.Depend(ModuleForCommon())

	return mb.Create()
}

func ModuleForServer() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theServerModuleResFS, theServerModuleResPath)

	mb.Components(server4mediapool.ExportComponents)

	mb.Depend(ModuleForCommon())
	mb.Depend(libgin.Module())
	mb.Depend(buckets.ModuleLib())

	mb.Depend(aliyun.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {
	mb := new(application.ModuleBuilder)

	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)

	mb.Components(test4mediapool.ExportComponents)

	mb.Depend(ModuleForClient())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
