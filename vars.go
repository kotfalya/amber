package db

const (
	RecordLevelMemory = iota
	RecordLevelDisk

	NetLevelNone = iota
	NetLevelAsync
	NetLevelSync
)

const (
	ReadLevelOptimistic = iota
	ReadLevelStable
)

const (
	WriteLevelOptimistic = iota
	WriteLevelTransactionApproved
	WriteLevelSaved
	WriteLevelPersist
)
