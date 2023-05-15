package helper

import (
	"crypto/rand"
	"math/big"
)

func RandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, 9)
	for i := range b {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[r.Int64()]
	}
	return string(b)
}
