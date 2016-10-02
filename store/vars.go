package store

import "flag"

const (
	ErrUndefinedKey   = "store:undefined key"
	ErrUnknowMessage  = "store:unknown message"
	ErrInvalidKeyType = "store:invalid key type"
)

var (
	pageKeysSize         = flag.Uint("pageKeysSize", 1000, "Max keys in one page")
	pageExpandTreshold   = flag.Uint("pageExpandTreshold", 90, "Threshold for a page expand (in percents). Should be between 1 and 100")
	pageCollapseTreshold = flag.Uint("pageCollapseTreshold", 50, "Threshold for a page collapse (in percents). Should be between 1 and 100")
	pageLeafPoolSize     = flag.Int("pageLeafPoolSize", 50, " Page leaf pool size")
	pageReqBufferSize    = flag.Int("pageReqBufferSize", 10, "Requests buffer size")
	checkPageSizeEvery   = flag.Int("checkPageSizeEvery", 1, "How often page will check own size to begin expand process (is seconds).")

	expandStartSize   = calculateExpandStartSize()
	collapseStartSize = calculateCollapseStartSize()
)

func calculateExpandStartSize() uint {
	return *pageKeysSize * *pageExpandTreshold / 100
}

func calculateCollapseStartSize() uint {
	return *pageKeysSize * *pageCollapseTreshold / 100
}
