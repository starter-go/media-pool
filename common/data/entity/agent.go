package entity

import "github.com/starter-go/libgorm"

type DBAgent interface {
	libgorm.Agent
}
