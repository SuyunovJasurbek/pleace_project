package service

import (
	"stad_projekt/models"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateData(data models.AparatData) (string, error) {

	var md = models.AparatDataToDB{
		Id:                uuid.NewString(),
		FeildId:           data.FeildId,
		ResultHumidity:    data.ResultHumidity,
		ResultTemperature: data.ResultTemperature,
		ResultLight:       data.ResultLight,
		CreateAt:          time.Now().Format("2006-01-02 15:04:05"),
	}

	id, err := s.repo.Admin().CreateData(md)
	if err != nil {
		return "", err
	}

	return id, nil

}

func (s *Service) GetData(data string ) ([]models.AparatDataToDB, error) {
	
	res, err := s.repo.Admin().GetData(data)
	if err != nil {
		return nil, err
	}

	return res, nil

}
