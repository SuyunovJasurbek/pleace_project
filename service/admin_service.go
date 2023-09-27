package service

import (
	"errors"
	"stad_projekt/models"
	"time"

	"github.com/google/uuid"
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

func (s *Service) CreateCountry (data models.Country) (string , error) {
	var cnt= models.CountryToDB{
		Id:       uuid.NewString(),
		Name:     data.Name,
		Location: data.Location,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	id , err :=s.repo.Admin().Country(cnt)
	if err != nil {
		 return "", err
	}
	return id, nil
}

func (s *Service) CreateField (data models.Feild)(string, error) {
	var md = models.FeildToDB{
		Id:       uuid.NewString(),
		CountryId: data.CountryId,
		Name:     data.Name,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	id , err :=s.repo.Admin().Feild(md)
	if err!= nil {
		return "", err
	}

	return id, nil
}

func (s *Service) CreatePicture (data models.Picture)(string, error) {
	var md = models.PictureToDB{
		Id:       uuid.NewString(),
		FeildId:  data.FeildId,
		Url:      data.Url,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	id , err :=s.repo.Admin().Picture(md)
	if err!= nil {
		return "", err
	}

	return id, nil
}

func (s *Service) GetCountry () ([]models.GetCountry, error) {
	return s.repo.Admin().GetCountry()
}

func (s *Service) GetField () ([]models.GetCountry, error) {
	return s.repo.Admin().GetCountry()
}


func (s *Service) GetPicture(dat models.GetFeildId) ([]models.GetPicture, error) {
	pictures, err := s.repo.Admin().GetPicture(dat)
	if err != nil {
		return nil, err
	}
	return pictures, nil
}