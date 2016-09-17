package utils

import "hash/adler32"

// Hash func determines child Page to use
func Hash(keyName string, pageChildSize int) (rv uint32, err error) {
	var checkSum uint32
	err = nil

	byteKeyName := []byte(keyName)
	checkSum = adler32.Checksum(byteKeyName)
	rv = checkSum % uint32(pageChildSize)
	return
}
