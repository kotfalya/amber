package db

import (
	"strings"
)

const (
	ErrInvalidResType = "cmd:invalid res type"

	OptionNotFound = -1

	RecordLevelTitle           = "level"
	RecordLevelOptimisticTitle = "optimistic"
	RecordLevelStableTitle     = "stable"
)

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
)

func ParseReadLevel(options []string) int {
	if len(options) == 0 {
		return OptionNotFound
	}

	var level string

	for i, v := range options {
		// check if next element exists and read it
		if strings.ToLower(v) == RecordLevelTitle && len(options) > i+1 {
			level = strings.ToLower(options[i+1])
			break
		}
	}

	if level == "" {
		return OptionNotFound
	}

	switch level {
	case RecordLevelOptimisticTitle:
		return ReadLevelOptimistic
	case RecordLevelStableTitle:
		return ReadLevelStable
	default:
		return OptionNotFound
	}
}
