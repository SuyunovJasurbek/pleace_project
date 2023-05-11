package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// f := "jasurbeksuyunov"
	// d, _ := GenerateHash(f)
	// fmt.Println(d)
	// var w string
	// fmt.Println(" parolni kiriting :")
	// fmt.Scan(&w)
	// r := CheckPasswordIfMatchs("$2a$10$I2Yjq0Vt/rRT2bDT5OFpru.76ub2jTqDWvna0N.dTHvTlyJEaaP0i", w)
	// fmt.Println(".....", r)
	f :=GenerateToken("4c0f9a3a-46d8-4038-91ba-54e7d3c776b0")
	fmt.Println(f)
}


func GenerateHash(password string) (string, error) {
	password_bytes := []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(password_bytes, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

// Comparing the password with the hash
func CheckPasswordIfMatchs(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
func GenerateToken(user_id string) string {
	secret_key := "AssalomuAlaykumXushKelibsizUstoz"
	tim := time.Now().String()
	x := secret_key + tim + user_id + "J"
	h := sha256.New()
	h.Write([]byte(x))
	hash := h.Sum(nil)
	authTokenStep := fmt.Sprintf("%x", hash)
	return authTokenStep
}
