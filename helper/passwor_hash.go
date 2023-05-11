package helper

import (
	"golang.org/x/crypto/bcrypt"
)
// Generate password hash
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