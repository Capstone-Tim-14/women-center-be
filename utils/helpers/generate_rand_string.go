package helpers

import (
	"math/rand"
	"time"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func StringWithCharset(length int) string {
	b := make([]rune, length)

	for i := range b {
		b[i] = charset[randomizer.Intn(len(charset))]
	}

	return string(b)
}
