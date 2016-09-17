package store

import "flag"

const (
	ErrUndefinedKey = "store:undefined key"
)

var (
	pageKeysSize         = flag.Int("pageKeysSize", 1000, "Max keys in one page")
	pageExpandTreshold   = flag.Float64("pageExpandTreshold", 0.90, "Treshold for a page expands")
	pageCollapseTreshold = flag.Float64("pageCollapseTreshold", 0.50, "Treshold for a page collapse")
	pageChildSize        = flag.Int("pageChildSize", 50, "size of children pool")
)
