package service

import (
	"fmt"
	"stad_projekt/models"
)

func (s *Service) SignUpPeraon(data models.SignInPersonModel) error {
	fmt.Println(data)
	err := s.repo.Person().SignUpPerson(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *Service) GetHumidity(device_id string) ([]models.GetHumidity, error) {
	fmt.Println("service", device_id)
	data, err := s.repo.Person().GetHumidity(device_id)
	if err != nil {
		fmt.Println(err)
		return []models.GetHumidity{}, err
	}
	return data, nil
}

func (s *Service) GetTemperature(device_id string) ([]models.GetTemperature, error) {
	fmt.Println("service", device_id)
	data, err :=s.repo.Person().GetTemperature(device_id)
	if err != nil {
		fmt.Println(err)
		return []models.GetTemperature{}, err
	}
	return data, nil
}

func (s *Service) GetLight(device_id string) ([]models.GetLight, error) {
	fmt.Println("service", device_id)
	data, err := s.repo.Person().GetLight(device_id)
	if err != nil {
		fmt.Println(err)
		return []models.GetLight{}, err
	}
	return data, nil
}

func (s *Service) GetHome(id string) (models.PersonSignInModel, error) {
	fmt.Println("service", id)
	return s.repo.Person().GetHome(id)
}

func (s *Service) SignInPerson(parol string) (string, error) {
	fmt.Println("service", parol)
	return s.repo.Person().SignInPerson(parol)
}
