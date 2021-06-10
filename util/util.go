package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random values between min and max
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string value
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		s := alphabets[rand.Intn(k)]
		sb.WriteByte(s)
	}
	return sb.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomMoney() int64 {
	return randomInt(0, 1000)
}

func RandomCurrency() string {
	c := []string{
		"USD",
		"EUR",
		"CAD",
		"INR",
	}
	n := len(c)
	return c[rand.Intn(n)]
}
