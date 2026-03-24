package dto

import "github.com/starter-go/media-pool/common/data/dxo"

type Example struct {
	ID dxo.ExampleID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
