package utils

import "hash/adler32"

// TextToIndex func determines scaled int hash by text
func TextToIndex(text string, scale int) int {
	byteText := []byte(text)
	checkSum := adler32.Checksum(byteText)

	return int((checkSum % uint32(scale)) - uint32(1))
}
