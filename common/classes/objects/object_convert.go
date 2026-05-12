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

	sum := src.Sum
	strSum := sum.String()
	date := src.CreatedAt

	dst.ID = src.ID
	dst.Name = src.Name
	dst.Type = src.Type
	dst.Length = src.Size
	dst.Sum = lang.Hex(strSum)
	dst.Type = src.Type
	dst.Location = src.Location

	dst.CreatedAt = date
	dst.UpdatedAt = date
	dst.Date = date.Time()

	return nil
}
