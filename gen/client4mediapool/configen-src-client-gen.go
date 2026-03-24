package client4mediapool
import (
    p700d48835 "github.com/starter-go/media-pool/client"
    p46220f67d "github.com/starter-go/media-pool/common/classes/objects"
     "github.com/starter-go/application"
)

// type p700d48835.MPClient in package:github.com/starter-go/media-pool/client
//
// id:com-700d48835613eea6-client-MPClient
// class:
// alias:alias-46220f67d06e6dbd28c3603d4b14f6ae-Client
// scope:singleton
//
type p700d488356_client_MPClient struct {
}

func (inst* p700d488356_client_MPClient) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-700d48835613eea6-client-MPClient"
	r.Classes = ""
	r.Aliases = "alias-46220f67d06e6dbd28c3603d4b14f6ae-Client"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p700d488356_client_MPClient) new() any {
    return &p700d48835.MPClient{}
}

func (inst* p700d488356_client_MPClient) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p700d48835.MPClient)
	nop(ie, com)

	
    com.FilterList = inst.getFilterList(ie)


    return nil
}


func (inst*p700d488356_client_MPClient) getFilterList(ie application.InjectionExt)[]p46220f67d.FilterRegistry{
    dst := make([]p46220f67d.FilterRegistry, 0)
    src := ie.ListComponents(".class-46220f67d06e6dbd28c3603d4b14f6ae-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p46220f67d.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}


