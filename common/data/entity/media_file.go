package entity

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/data/dxo"
)

type MediaFile struct {
	ID dxo.MediaFileID

	Base

	SimpleName string

	ContentType string

	ContentLength int64

	Path objects.Path

	Sum lang.Hex `gorm:"unique"` // the sha256 sum
}
