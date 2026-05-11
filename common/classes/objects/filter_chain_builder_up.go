package objects

import "sort"

////////////////////////////////////////////////////////////////////////////////

type UploadFilterChainBuilder struct {
	filters []*FilterRegistration
}

func (inst *UploadFilterChainBuilder) Build() UploadFilterChain {

	var chain UploadFilterChain
	chain = new(innerUploadFilterChainEnding)

	all := inst.filters
	all = inst.innerPrepareList(all)

	for _, item := range all {

		if !inst.innerIsReady(item) {
			continue
		}

		node := new(innerUploadFilterChainNode)
		node.next = chain
		node.f = item.Up
		chain = node
	}

	return chain
}

func (inst *UploadFilterChainBuilder) AddFilter(f UploadFilter) {

	if f == nil {
		return
	}

	reg := &FilterRegistration{
		Enabled:  true,
		Priority: 0,
		Label:    "unnamed",
		Up:       f,
		Down:     nil,
	}

	inst.filters = append(inst.filters, reg)
}

func (inst *UploadFilterChainBuilder) AddRegistration(reg *FilterRegistration) {

	if !inst.innerIsReady(reg) {
		return
	}

	inst.filters = append(inst.filters, reg)
}

func (inst *UploadFilterChainBuilder) AddRegistry(reg FilterRegistry) {

	if reg == nil {
		return
	}

	src := reg.ListFilters()
	inst.filters = append(inst.filters, src...)
}

func (inst *UploadFilterChainBuilder) innerIsReady(reg *FilterRegistration) bool {

	if reg == nil {
		return false
	}

	if !reg.Enabled {
		return false
	}

	if reg.Up == nil {
		return false
	}

	return true
}

func (inst *UploadFilterChainBuilder) innerPrepareList(src []*FilterRegistration) []*FilterRegistration {

	dst := make([]*FilterRegistration, 0)

	for _, item := range src {
		if !inst.innerIsReady(item) {
			continue
		}
		dst = append(dst, item)
	}

	sor := new(innerFilterSorter)
	return sor.sort(dst)
}

////////////////////////////////////////////////////////////////////////////////

type innerFilterSorter struct {
	filters []*FilterRegistration
	// reverse bool
}

func (inst *innerFilterSorter) sort(src []*FilterRegistration) []*FilterRegistration {

	if src == nil {
		return make([]*FilterRegistration, 0)
	}

	inst.filters = src
	// inst.reverse = reverse

	sort.Sort(inst)
	return inst.filters
}

func (inst *innerFilterSorter) numberOf(index int) int {
	o := inst.filters[index]
	if o == nil {
		return 0xfffffff
	}
	return o.Priority
}

func (inst *innerFilterSorter) Len() int {
	return len(inst.filters)
}
func (inst *innerFilterSorter) Less(i1, i2 int) bool {
	n1 := inst.numberOf(i1)
	n2 := inst.numberOf(i2)

	// if inst.reverse {
	// 	return (n1 > n2)
	// }

	return (n1 < n2)
}
func (inst *innerFilterSorter) Swap(i1, i2 int) {
	list := inst.filters
	list[i1], list[i2] = list[i2], list[i1]
}

////////////////////////////////////////////////////////////////////////////////

type innerUploadFilterChainNode struct {
	f    UploadFilter
	next UploadFilterChain
}

// Put implements UploadFilterChain.
func (i *innerUploadFilterChainNode) Put(o *IOContext) error {
	f := i.f
	n := i.next
	return f.Put(o, n)
}

////////////////////////////////////////////////////////////////////////////////

type innerUploadFilterChainEnding struct {
	// NONE
}

// Put implements UploadFilterChain.
func (i *innerUploadFilterChainEnding) Put(o *IOContext) error {
	// NOP
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
