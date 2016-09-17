package store

import "flag"

var (
	pageKeysSize  = flag.Int("pageKeysSize", 1000, "Max keys in one page")
	pageChildSize = flag.Int("pageChildSize", 50, "size of children pool")
)
