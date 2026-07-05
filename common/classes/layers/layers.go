package layers

const (
	PriorityMin = iota

	PriorityStorage
	PriorityDB
	PriorityCache
	PriorityMeta

	PriorityURL
	PriorityFiles
	PriorityPath
	PrioritySum

	PriorityDebug
	PriorityMonitor
	PriorityParamsChecker
	PriorityBuffer
	PriorityTemp

	PriorityOther
	PriorityMax
)
