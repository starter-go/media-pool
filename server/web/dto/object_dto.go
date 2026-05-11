package dto

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/common/data/dxo"
)

type Object struct {

	// id
	ID dxo.ObjectID `json:"id"`

	// base
	Base

	// fields

	Name string `json:"name"`

	Type string `json:"type"`

	Sum lang.Hex `json:"sha256sum"`

	Length int64 `json:"length"`
}
