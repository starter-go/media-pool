package objects

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/media-pool/server/web/dto"
)

func Convert2DTO(src *Object, dst *dto.Object) error {

	if src == nil || dst == nil {
		return fmt.Errorf("param: src|dst is nil")
	}

	meta := &src.Meta
	sum := meta.Sum
	strSum := sum.String()

	dst.ID = meta.ID
	dst.Name = meta.Name
	dst.Type = meta.Type
	dst.Length = meta.Length
	dst.Sum = lang.Hex(strSum)
	dst.Type = meta.Type
	dst.Location = src.Location

	dst.CreatedAt = meta.CreatedAt
	dst.UpdatedAt = meta.UpdatedAt
	dst.Date = meta.UpdatedAt.Time()

	return nil
}
