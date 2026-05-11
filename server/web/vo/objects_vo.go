package vo

import "github.com/starter-go/media-pool/server/web/dto"

type Objects struct {

	// base
	Base

	Items []*dto.Object `json:"objects"`
}
