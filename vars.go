package db

import (
	"errors"
	"flag"
	"strings"
)

var (
	pageKeysSize          = flag.Uint("pageKeysSize", 1000, "Max keys in one page")
	pageExpandThreshold   = flag.Uint("pageExpandThreshold", 90, "Threshold for a page expand (in percents). Should be between 1 and 100")
	pageCollapseThreshold = flag.Uint("pageCollapseThreshold", 50, "Threshold for a page collapse (in percents). Should be between 1 and 100")
	pageLeafPoolSize      = flag.Uint("pageLeafPoolSize", 50, " Page leaf pool size")
	pageReqBufferSize     = flag.Int("pageReqBufferSize", 10, "Requests buffer size")
	checkPageSizeEvery    = flag.Int("checkPageSizeEvery", 1, "How often page will check own size to begin expand process (is seconds).")

	expandStartSize   = calculateExpandStartSize()
	collapseStartSize = calculateCollapseStartSize()
)

func calculateExpandStartSize() uint {
	return *pageKeysSize * *pageExpandThreshold / 100
}

func calculateCollapseStartSize() uint {
	return *pageKeysSize * *pageCollapseThreshold / 100
}

const (
	// TODO add args to error
	ErrUndefinedKey      = "db: undefined key"
	ErrInvalidKeyType    = "db: invalid key type"
	ErrUndefinedKeyCmd   = "db: undefined key command"
	ErrInvalidResType    = "db: invalid res type"
	ErrInvalidReqHandler = "db: invalid req handler"
	ErrStopRequest       = "db: req was interrupted"

	ErrUnknownErrorCode = "db: unknown error code"
	ErrOptionNotFound   = "db: option not found"
	ErrOptionInvalid    = "db: option is invalid"

	LevelTitle           = "level"
	LevelOptimisticTitle = "optimistic"
	LevelLocalTitle      = "local"
	LevelMasterTitle     = "master"

	PersistTitle      = "persist"
	PersistNoneTitle  = "none"
	PersistAsyncTitle = "async"
	PersistSyncTitle  = "sync"
)

const (
	KeyCmdModeRead = iota
	KeyCmdModeUpdate
	KeyCmdModeUpsert
)

const (
	OptionNotFound = -1 * iota
	OptionInvalid
)

const (
	PersistNone = iota
	PersistAsync
	PersistSync
)

const (
	LevelOptimistic = iota
	LevelLocal
	LevelMaster
)

func readOption(option, defaultOption int) (int, error) {
	if option >= 0 {
		return option, nil
	} else {
		return defaultOption, parseErrorCode(option)
	}
}

func parseErrorCode(code int) error {
	switch code {
	case OptionInvalid:
		return errors.New(ErrOptionInvalid)
	case OptionNotFound:
		return errors.New(ErrOptionNotFound)
	default:
		panic(ErrUnknownErrorCode)
	}
}

func parseOption(optionName string, options []string) string {
	if len(options) == 0 {
		return ""
	}

	for i, v := range options {
		// check if next element exists and read it
		if strings.ToLower(v) == optionName && len(options) > i+1 {
			return strings.ToLower(options[i+1])
		}
	}

	return ""
}

func parseLevel(options []string) int {
	if level := parseOption(LevelTitle, options); level == "" {
		return OptionNotFound
	} else {
		switch level {
		case LevelOptimisticTitle:
			return LevelOptimistic
		case LevelLocalTitle:
			return LevelLocal
		case LevelMasterTitle:
			return LevelMaster
		default:
			return OptionInvalid
		}
	}
}

func parsePersist(options []string) int {
	if persist := parseOption(PersistTitle, options); persist == "" {
		return OptionNotFound
	} else {
		switch persist {
		case PersistNoneTitle:
			return PersistNone
		case PersistAsyncTitle:
			return PersistAsync
		case PersistSyncTitle:
			return PersistSync
		default:
			return OptionInvalid
		}
	}
}
