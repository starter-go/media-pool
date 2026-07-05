package imediafiles

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/common/classes/mediafiles"
	"github.com/starter-go/media-pool/common/data/entity"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

type MediaFileDaoImpl struct {

	//starter:component

	_as func(mediafiles.DAO) //starter:as("#")

	Agent   entity.DBAgent     //starter:inject("#")
	UUIDGen random.UUIDService //starter:inject("#")

}

// Insert implements mediafiles.DAO.
func (inst *MediaFileDaoImpl) Insert(db *gorm.DB, item *mediafiles.Entity) (*mediafiles.Entity, error) {

	item.ID = 0
	item.UUID = inst.innerGenUUID()

	db = inst.Agent.DB(db)
	res := db.Create(item)
	err := res.Error

	return item, err
}

// Find implements mediafiles.DAO.
func (inst *MediaFileDaoImpl) Find(db *gorm.DB, id mediafiles.ID) (*mediafiles.Entity, error) {

	db = inst.Agent.DB(db)
	item := inst.innerMakeItem()

	res := db.First(item, id)
	err := res.Error

	return item, err
}

func (inst *MediaFileDaoImpl) innerGenUUID() lang.UUID {

	ser := inst.UUIDGen
	b := ser.Build()
	b.Class("mediafiles.Entity")

	return b.Generate()
}

func (inst *MediaFileDaoImpl) innerMakeItem() *mediafiles.Entity {
	return new(mediafiles.Entity)
}

func (inst *MediaFileDaoImpl) innerMakeItemList() []*mediafiles.Entity {
	return make([]*mediafiles.Entity, 0)
}

func (inst *MediaFileDaoImpl) _impl() mediafiles.DAO {
	return inst
}
