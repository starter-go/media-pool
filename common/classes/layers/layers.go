package layers

const (
	PriorityMin = iota

	PriorityPool
	PriorityCache
	PriorityMonitor

	PriorityPath
	PrioritySum
	PriorityBuffer
	PriorityTemp
	PriorityDebug

	PriorityOther
	PriorityMax
)
