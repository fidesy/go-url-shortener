package shortener

import (
	"math/rand"
	"strings"
)

var (
	signs = "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetRandomSequence(length int) string {
	// Returns random string
	var str = strings.Builder{}
	str.Grow(length)

	for i := 0; i < length; i++ {
		str.WriteString(string(signs[rand.Intn(len(signs))]))
	}

	return str.String()
}
