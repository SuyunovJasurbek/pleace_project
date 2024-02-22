package service

import (
	"errors"
	"stad_projekt/models"
)

func (s *Service) SignIn(sgn models.SignInModel) (string ,  error) {

	token , err :=s.repo.Admin().SignIn(sgn)
	if err!= nil {
		return "", errors.New("error passport or login")
	}
	return token,nil
}

func (s *Service) Auth(token string) (bool) {
	return s.repo.Admin().Auth(token)
}


func (s *Service) GetInactiveUsers () ([]models.AccsesUser , error) {
	return s.repo.Admin().GetInactiveUsers()
}

func (s *Service) GetActiveUsers () ([]models.AccsesUser , error) {
	return s.repo.Admin().GetActiveUsers()
}

func (s *Service) CreatePersonCountry(data models.PersonCountry) error {
	return s.repo.Admin().CreatePersonCountry(data)
}

func (s * Service)  GetActivePleaces (person_id string ) ([]models.Place , error) {
	return s.repo.Admin().GetActivePleaces(person_id)
}

func (s *Service) UpdatePleace(pleace_id string) error {
	return s.repo.Admin().UpdatePleace(pleace_id)
}
func (s *Service) UpdatePerson(person_id string) error {
	return s.repo.Admin().UpdatePerson(person_id)
}
