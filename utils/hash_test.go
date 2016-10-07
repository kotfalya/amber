package utils

import (
	"fmt"
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
	etalon := iterations / scale
	m := make(map[uint]uint32)

	for i = 0; i < iterations; i++ {
		index := GetIndex(RandomString(10), scale, 12345)
		m[index] += 1

	}
	fmt.Println(m)

	for _, v := range m {
		if v >= etalon {
			if x := (v - (iterations / scale)); x > etalon/100*3 {
				fmt.Println(v - (iterations / scale))
			}
		} else {
			if x := ((iterations / scale) - v); x > etalon/100*3 {
				fmt.Println((iterations / scale) - v)
			}
		}
	}
}
