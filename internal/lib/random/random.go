package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int64) string {
	var (
		charset               = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	)

	res := make([]byte, size)
	for i := range res {
		res[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(res)
}
