package service

import (
	"errors"
	"fmt"
	"stad_projekt/helper"
	"stad_projekt/models"
)

func (s *Service) SignIn(sgn models.SignInModel) (models.SignInHandlerResponse, error) {
	//1.
	resp, err := s.s.Admin().SignIn(sgn.Login)
	if err != nil {
		return models.SignInHandlerResponse{}, errors.New("login yoki parol xato")
	}

	//2.
	fmt.Println("--------------------------------")
	fmt.Println(resp)
	fmt.Println("--------------------------------")

	if helper.CheckPasswordIfMatchs(resp.PasswordHash, sgn.Password) {
		token := helper.GenerateToken(resp.ID)
		var res = models.SignInHandlerResponse{
			ID:    resp.ID,
			Token: token,
		}
		return res, nil
	} else {
		return models.SignInHandlerResponse{}, errors.New("login yoki parol xato")
	}
}