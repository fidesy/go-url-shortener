package shortener

import (
	"math/rand"
)

var (
	signs = "0123456789abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetRandomSequence(length int) string {
	// Returns random string
	var str string
	for len(str) < length {
		str += string(signs[rand.Intn(len(signs))])
	}

	return str
}
