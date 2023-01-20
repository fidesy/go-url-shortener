package shortener

import (
	"math/rand"
	"strings"
	"time"
)

var (
	signs = "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Returns random string
func GetRandomSequence(length int) string {
	// seed for generating pseudo random numbers
	rand.Seed(time.Now().UnixNano())

	var str = strings.Builder{}
	str.Grow(length)

	for i := 0; i < length; i++ {
		str.WriteString(string(signs[rand.Intn(len(signs))]))
	}

	return str.String()
}
