package common4mediapool
import (
    p2fe7d86b0 "github.com/starter-go/media-pool/common/impl/icom"
     "github.com/starter-go/application"
)

// type p2fe7d86b0.Example in package:github.com/starter-go/media-pool/common/impl/icom
//
// id:com-2fe7d86b0cbb8f58-icom-Example
// class:
// alias:
// scope:singleton
//
type p2fe7d86b0c_icom_Example struct {
}

func (inst* p2fe7d86b0c_icom_Example) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2fe7d86b0cbb8f58-icom-Example"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2fe7d86b0c_icom_Example) new() any {
    return &p2fe7d86b0.Example{}
}

func (inst* p2fe7d86b0c_icom_Example) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2fe7d86b0.Example)
	nop(ie, com)

	


    return nil
}


