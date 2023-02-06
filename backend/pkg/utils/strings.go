package utils

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewRandomString(digit int) string {
	rand.Seed(time.Now().UnixNano())
	result := ""
	for i := 0; i < digit; i++ {
		result += string(charset[int(rand.Int())%len(charset)])
	}

	return result
}

func NewULID() string {
	return ulid.Make().String()
}
