package test4mediapool
import (
    p5df4e4e1c "github.com/starter-go/media-pool/src/test/golang/unittestcases"
     "github.com/starter-go/application"
)

// type p5df4e4e1c.ExampleCase in package:github.com/starter-go/media-pool/src/test/golang/unittestcases
//
// id:com-5df4e4e1c8ef0b1b-unittestcases-ExampleCase
// class:
// alias:
// scope:singleton
//
type p5df4e4e1c8_unittestcases_ExampleCase struct {
}

func (inst* p5df4e4e1c8_unittestcases_ExampleCase) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-5df4e4e1c8ef0b1b-unittestcases-ExampleCase"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p5df4e4e1c8_unittestcases_ExampleCase) new() any {
    return &p5df4e4e1c.ExampleCase{}
}

func (inst* p5df4e4e1c8_unittestcases_ExampleCase) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p5df4e4e1c.ExampleCase)
	nop(ie, com)

	


    return nil
}


