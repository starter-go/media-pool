package layers

const (
	PriorityMin = iota

	PriorityStorage
	PriorityCache
	PriorityMonitor

	PriorityPath
	PrioritySum
	PriorityBuffer
	PriorityTemp
	PriorityDebug
	PriorityParamsChecker

	PriorityOther
	PriorityMax
)
