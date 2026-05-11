package objects

import "context"

type IOContext struct {
	CC   context.Context
	Want *Object
	Have *Object
}
