package utils

import "hash/adler32"

// TextToIndex func determines scaled int hash by text
func TextToIndex(text string, scale int) uint32 {
	byteText := []byte(text)
	checkSum := adler32.Checksum(byteText)

	return checkSum % uint32(scale)
}
