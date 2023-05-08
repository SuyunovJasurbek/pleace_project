package service

import (
	"errors"
	"fmt"
	"stad_projekt/models"
)


func (s *Service) SignIn(sgn models.SignInModel) (string , error) {
	//1.
	resp, err :=s.s.Admin().SignIn(sgn.Login)
	if err!= nil {
		return "", errors.New("login yoki parol xato")
	}
	//2.
	fmt.Println("--------------------------------")
	fmt.Println(resp)
	fmt.Println("--------------------------------")
	return "", nil
}