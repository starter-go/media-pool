package idb

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/media-pool/common/data/entity"
)

type MPoolDataGroup struct {

	//starter:component

	_as func(libgorm.GroupRegistry) //starter:as(".")

	Prefix  string //starter:inject("${datagroup.mpool.prefix}")
	URI     string //starter:inject("${datagroup.mpool.uri}")
	Alias   string //starter:inject("${datagroup.mpool.alias}")
	Source  string //starter:inject("${datagroup.mpool.source}")
	Enabled bool   //starter:inject("${datagroup.mpool.enabled}")

}

// Prototypes implements libgorm.Group.
func (inst *MPoolDataGroup) Prototypes() []any {
	prefix := inst.Prefix
	all := entity.ListEntities(prefix)
	return all
}

// Groups implements libgorm.GroupRegistry.
func (inst *MPoolDataGroup) Groups() []*libgorm.GroupRegistration {

	g1 := &libgorm.GroupRegistration{
		Group: inst,

		Enabled: inst.Enabled,
		Alias:   inst.Alias,
		URI:     inst.URI,
		Prefix:  inst.Prefix,
		Source:  inst.Source,
	}

	list := make([]*libgorm.GroupRegistration, 0)
	list = append(list, g1)
	return list
}

func (inst *MPoolDataGroup) _impl() (libgorm.GroupRegistry, libgorm.Group) {
	return inst, inst
}
