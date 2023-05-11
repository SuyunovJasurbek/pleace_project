package helper

import (

	"fmt"

	"golang.org/x/crypto/bcrypt"
	
)

func GenerateHash(password string ) (string , error) {

	password_bytes:=[]byte(password)

	hashedPasswordBytes, err :=bcrypt.GenerateFromPassword(password_bytes,bcrypt.DefaultCost)

	if err!= nil {

		return "", err

	}

	fmt.Println("------------------------------------")

	fmt.Println("HashedPassword :",hashedPasswordBytes)

	fmt.Println("------------------------------------")

	return string(hashedPasswordBytes), nil
}


// Comparing the password with the hash
func CheckPasswordIfMatchs(hashedPassword, currPassword string) bool {

	err := bcrypt.CompareHashAndPassword(

		[]byte(hashedPassword), []byte(currPassword))

	return err == nil

}