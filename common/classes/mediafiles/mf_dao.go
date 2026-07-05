package mediafiles

import "gorm.io/gorm"

type DAO interface {
	Find(db *gorm.DB, id ID) (*Entity, error)

	Insert(db *gorm.DB, item *Entity) (*Entity, error)
}
