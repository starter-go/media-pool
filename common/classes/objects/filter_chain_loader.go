package objects

type FilterChainLoader struct {
	filters []*FilterRegistration
}

func (inst *FilterChainLoader) Init(list []FilterRegistry) {

	inst.filters = nil

	src := list
	dst := inst.filters

	for _, r1 := range src {
		tmp := r1.ListFilters()
		for _, r2 := range tmp {
			dst = append(dst, r2)
		}
	}

	inst.filters = dst
}

func (inst *FilterChainLoader) LoadChainUp() UploadFilterChain {

	builder := new(UploadFilterChainBuilder)
	list := inst.filters

	for _, r := range list {
		builder.AddRegistration(r)
	}

	return builder.Build()
}

func (inst *FilterChainLoader) LoadChainDown() DownloadFilterChain {

	builder := new(DownloadFilterChainBuilder)
	list := inst.filters

	for _, r := range list {
		builder.AddRegistration(r)
	}

	return builder.Build()
}
