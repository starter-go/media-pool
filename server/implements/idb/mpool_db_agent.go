package idb

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/media-pool/common/data/entity"
	"gorm.io/gorm"
)

type MPoolDBAgentImpl struct {

	//starter:component

	_as func(entity.DBAgent) //starter:as("#")

	DSM libgorm.DataSourceManager //starter:inject("#")

	Source string //starter:inject("${datagroup.mpool.source}")

	cached *gorm.DB
}

// DB implements entity.DBAgent.
func (inst *MPoolDBAgentImpl) DB(db *gorm.DB) *gorm.DB {

	if db != nil {
		return db
	}

	db = inst.cached
	if db != nil {
		return db
	}

	db, err := inst.innerGetDB()
	if err != nil {
		panic(err)
	}

	return db
}

func (inst *MPoolDBAgentImpl) innerGetDB() (*gorm.DB, error) {

	db := inst.cached
	if db != nil {
		return db, nil
	}

	db, err := inst.innerLoadDB()
	if err != nil {
		return nil, err
	}

	inst.cached = db
	return db, nil
}

func (inst *MPoolDBAgentImpl) innerLoadDB() (*gorm.DB, error) {
	alias := inst.Source
	src, err := inst.DSM.GetDataSource(alias)
	if err != nil {
		return nil, err
	}
	return src.DB()
}

func (inst *MPoolDBAgentImpl) _impl() entity.DBAgent {
	return inst
}
