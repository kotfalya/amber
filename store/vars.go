package store

import "flag"

const (
	ErrUndefinedKey   = "store:undefined key"
	ErrUnknowMessage  = "store:unknow message"
	ErrInvalidKeyType = "store: invalid key type"
)

var (
	pageKeysSize         = flag.Int("pageKeysSize", 1000, "Max keys in one page")
	pageExpandTreshold   = flag.Float64("pageExpandTreshold", 0.90, "Treshold for a page expands")
	pageCollapseTreshold = flag.Float64("pageCollapseTreshold", 0.50, "Treshold for a page collapse")
	pageLeafPoolSize     = flag.Int("pageLeafPoolSize", 50, " Page leaf pool size")
	pageReqBufferSize    = flag.Int("pageReqBufferSize", 10, "Requests buffer size")
)
