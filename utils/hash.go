package utils

import "github.com/OneOfOne/xxhash"

// GetIndex func determines scaled int hash by text
func GetIndex(text string, scale uint32, seed uint32) uint {
	byteText := []byte(text)
	hash := xxhash.Checksum32S(byteText, seed)

	return uint((hash % scale))
}
