package storage

import (
	"context"
	"stad_projekt/models"
)

type StorageI interface {
	Person() PersonI
	Admin() AdminI
	CloseDb()
}

type CasheStorageI interface {
	Token() TokenI
}
type PersonI interface {	
	SignUpPerson(models.SignInPersonModel) error
	SignInPerson(parol string) (string, error)
	GetHumidity(device_id string) ([]models.GetHumidity, error)
	GetTemperature(device_id string) ([]models.GetTemperature, error)
	GetLight(device_id string) ([]models.GetLight, error)
	GetDayDate(device_id string) ([]models.DayDate, error)
	GetHome(id string) (models.PersonSignInModel, error)
	SentToCountry(person_id string) error
}

type TokenI interface {
	SetToken(ctx context.Context, key, value string) error
	GetToken(ctx context.Context,key string) (token string, err error)
}

type AdminI interface {
	SignIn(models.SignInModel) (string, error)
	Auth(token string) bool
	CreateData(models.AparatDataToDB) (string, error)
	GetData(dat string) ([]models.AparatDataToDB, error)
	GetInactiveUsers() ([]models.AccsesUser, error)
	GetActiveUsers() ([]models.AccsesUser, error)
	GetActivePleaces(person_id string) ([]models.Place, error)
	CreatePersonCountry(models.PersonCountry) error
	UpdatePleace(pleace_id string) error
	UpdatePerson(person_id string) error
}