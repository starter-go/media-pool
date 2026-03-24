package vo

import "github.com/starter-go/media-pool/server/web/dto"

type Examples struct {
	Base

	Items []*dto.Example `json:"examples"`
}
