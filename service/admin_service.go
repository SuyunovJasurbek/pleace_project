package service

import (
	"errors"
	"stad_projekt/helper"
	"stad_projekt/models"

	"github.com/google/uuid"
)

func (s *Service) SignIn(sgn models.SignInModel) (models.SignInHandlerResponse, error) {
	//1.
	resp, err := s.repo.Admin().SignIn(sgn.Login)
	if err != nil {
		return models.SignInHandlerResponse{}, errors.New("login yoki parol xato")
	}

	//2.
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
func (s *Service) CreateStadium() (string, error) {
	return "", nil
}
func (s *Service) CreatePicture(user_id string) (string, error) {
	return "122334534534", nil
}
func (s *Service) CreateStadiumName(st models.CreateStadiumNameRequest) (string, error) {

	var std = models.StadiumNameRequest{
		ID:      uuid.NewString(),
		User_Id: st.User_Id,
		Name:    st.Name,
	}

	stadium_id, err := s.repo.Admin().CreateStadiumName(std)
	if err != nil {
		return "", errors.New("create name error db")
	}
	return stadium_id, nil
}
func (s *Service) CheckStadiumPicture(st models.CheckStadiumPicture) bool {
	return s.repo.Admin().CheckStadiumPicture(st)
}