package mediafiles

import "github.com/starter-go/security-gorm/rbacdb"

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)

	dst.ContentLength = src.ContentLength
	dst.ContentType = src.ContentType
	dst.Path = src.Path
	dst.SimpleName = src.SimpleName
	dst.Sum = src.Sum

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)

	dst.ContentLength = src.ContentLength
	dst.ContentType = src.ContentType
	dst.Path = src.Path
	dst.SimpleName = src.SimpleName
	dst.Sum = src.Sum

	return nil
}
