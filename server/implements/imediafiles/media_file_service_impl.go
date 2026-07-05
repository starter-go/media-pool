package imediafiles

import (
	"context"

	"github.com/starter-go/media-pool/common/classes/mediafiles"
)

type MediaFileServiceImpl struct {

	//starter:component

	_as func(mediafiles.Service) //starter:as("#")

	Dao mediafiles.DAO //starter:inject("#")

}

// Find implements mediafiles.Service.
func (inst *MediaFileServiceImpl) Find(cc context.Context, id mediafiles.ID) (*mediafiles.DTO, error) {

	it1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}

	it2 := new(mediafiles.DTO)
	err = mediafiles.ConvertE2D(it1, it2)
	return it2, err
}

// Insert implements mediafiles.Service.
func (inst *MediaFileServiceImpl) Insert(cc context.Context, item *mediafiles.DTO) (*mediafiles.DTO, error) {

	it2 := new(mediafiles.Entity)
	it4 := new(mediafiles.DTO)

	err := mediafiles.ConvertD2E(item, it2)
	if err != nil {
		return nil, err
	}

	it3, err := inst.Dao.Insert(nil, it2)
	if err != nil {
		return nil, err
	}

	err = mediafiles.ConvertE2D(it3, it4)
	return it4, err
}

func (inst *MediaFileServiceImpl) _impl() mediafiles.Service {
	return inst
}
