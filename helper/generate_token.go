package helper

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GenerateToken(user_id string) (string) {
	secret_key := "AssalomuAlaykumXushKelibsizUstoz"
	tim := time.Now().String()
	x := secret_key + tim + user_id+"J"
	h := sha256.New()
	h.Write([]byte(x))
	hash := h.Sum(nil)
	authTokenStep := fmt.Sprintf("%x", hash)
	return authTokenStep
}