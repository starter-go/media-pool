package layers

const (
	PriorityMin = iota

	PriorityStorage
	PriorityCache
	PriorityMonitor

	PriorityPath
	PriorityURL
	PrioritySum
	PriorityBuffer
	PriorityTemp
	PriorityDebug
	PriorityParamsChecker

	PriorityOther
	PriorityMax
)
