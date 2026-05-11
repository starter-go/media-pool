package objects

////////////////////////////////////////////////////////////////////////////////

type DownloadFilterChainBuilder struct {
	filters []*FilterRegistration
}

func (inst *DownloadFilterChainBuilder) Build() DownloadFilterChain {

	var chain DownloadFilterChain
	chain = new(innerDownloadFilterChainEnding)

	all := inst.filters
	all = inst.innerPrepareList(all)

	for _, item := range all {

		if !inst.innerIsReady(item) {
			continue
		}

		node := new(innerDownloadFilterChainNode)
		node.next = chain
		node.f = item.Down
		chain = node
	}

	return chain
}

func (inst *DownloadFilterChainBuilder) AddFilter(f DownloadFilter) {

	if f == nil {
		return
	}

	reg := &FilterRegistration{
		Enabled:  true,
		Priority: 0,
		Label:    "unnamed",
		Up:       nil,
		Down:     f,
	}

	inst.filters = append(inst.filters, reg)
}

func (inst *DownloadFilterChainBuilder) AddRegistration(reg *FilterRegistration) {

	if !inst.innerIsReady(reg) {
		return
	}

	inst.filters = append(inst.filters, reg)
}

func (inst *DownloadFilterChainBuilder) AddRegistry(reg FilterRegistry) {

	if reg == nil {
		return
	}

	src := reg.ListFilters()
	inst.filters = append(inst.filters, src...)
}

func (inst *DownloadFilterChainBuilder) innerIsReady(reg *FilterRegistration) bool {

	if reg == nil {
		return false
	}

	if !reg.Enabled {
		return false
	}

	if reg.Down == nil {
		return false
	}

	return true
}

func (inst *DownloadFilterChainBuilder) innerPrepareList(src []*FilterRegistration) []*FilterRegistration {

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

type innerDownloadFilterChainNode struct {
	f    DownloadFilter
	next DownloadFilterChain
}

// Put implements DownloadFilterChain.
func (i *innerDownloadFilterChainNode) Fetch(o *IOContext) error {
	f := i.f
	n := i.next
	return f.Fetch(o, n)
}

////////////////////////////////////////////////////////////////////////////////

type innerDownloadFilterChainEnding struct {
	// NONE
}

// Put implements DownloadFilterChain.
func (i *innerDownloadFilterChainEnding) Fetch(o *IOContext) error {
	// NOP
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
