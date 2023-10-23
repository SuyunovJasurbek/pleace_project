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

func (s *Service) GetField () ([]models.GetField, error) {
	return s.repo.Admin().GetField()
}

func (s *Service) GetPicture(dat string) ([]models.GetPicture, error) {
	pictures, err := s.repo.Admin().GetPicture(dat)
	if err != nil {
		return nil, err
	}
	return pictures, nil
}

func (s *Service) UpdateCountry(data models.UpdateCountry) (string, error) {
	var dat= models.CountryToDB{

		Id:       data.Id,
		Name:     data.Name,
		Location: data.Location,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	_,err :=s.repo.Admin().UpdateCountry(dat)
	if err != nil {
		return "", err
	}
	return "ok", nil
}

func (s *Service) UpdateField(data models.FeildToDB) (string, error) {
	return s.repo.Admin().UpdateField(data)
}

func (s *Service) DeleteCountry(country_id string) (string, error) {
	return s.repo.Admin().DeleteCountry(country_id)
}

func (s *Service) DeleteField(field_id string) (string, error) {
	return s.repo.Admin().DeleteField(field_id)
}