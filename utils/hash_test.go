package utils

import (
	"math/rand"
	"testing"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func TestGetIndex(t *testing.T) {
	var i uint32
	var scale uint32 = 50
	var iterations uint32 = 1000000
	var acceptable_percent uint32 = 2

	etalon := iterations / scale
	limit := etalon / 100 * acceptable_percent
	m := make(map[uint]uint32)

	for i = 0; i < iterations; i++ {
		index := GetIndex(RandomString(10), scale, 12345)
		m[index] += 1
	}

	for _, v := range m {
		x := int(v) - int(iterations/scale)
		if x < 0 {
			x *= -1
		}
		if uint32(x) > limit {
			t.Error("Your balance - sucks!")
		}
	}
}
