package dto

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/common/classes/objects"
	"github.com/starter-go/media-pool/common/data/dxo"
	"github.com/starter-go/rbac"
)

type Base struct {
	rbac.BaseDTO
}

type MediaFile struct {
	ID dxo.MediaFileID

	Base

	SimpleName    string       `json:"name"`
	ContentLength int64        `json:"length"`
	ContentType   string       `json:"type"`
	Path          objects.Path `json:"path"`

	Sum lang.Hex `json:"sha256sum"`
}
