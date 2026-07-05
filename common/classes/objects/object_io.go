package objects

import "context"

////////////////////////////////////////////////////////////////////////////////

type InputOutputContext struct {
	CC   context.Context
	Want *Object
	Have *Object
}

// alias:IOC
type IOC = InputOutputContext

// alias:IOContext
type IOContext = InputOutputContext

////////////////////////////////////////////////////////////////////////////////

func (inst *InputOutputContext) GetHave(autoMake bool) *Object {
	o := inst.Have
	if o == nil && autoMake {
		o = new(Object)
		inst.Have = o
	}
	return o
}

func (inst *InputOutputContext) GetWant(autoMake bool) *Object {
	o := inst.Want
	if o == nil && autoMake {
		o = new(Object)
		inst.Want = o
	}
	return o
}

////////////////////////////////////////////////////////////////////////////////
// EOF
