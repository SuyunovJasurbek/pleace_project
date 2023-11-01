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
}

type TokenI interface {
	SetToken(ctx context.Context, key, value string) error
	GetToken(ctx context.Context,key string) (token string, err error)
}

type AdminI interface {
	SignIn(models.SignInModel) (string, error)
	Auth(token string) bool
	UserList(models.UserList) ([]models.UserList, error)
	Country(models.CountryToDB) (string, error)
	Feild(models.FeildToDB) (string, error)
	Picture(models.PictureToDB) (string, error)
	GetCountry() ([]models.GetCountry, error)
	GetField() ([]models.GetField, error)
	CreateData(models.AparatDataToDB) (string, error)
	GetData(dat string) ([]models.AparatDataToDB, error)
	GetPicture(dat string) ([]models.GetPicture, error)
	GetFeildIdToList(country_id string) ([]models.CountryToDB, error)
	DeleteCountry(country_id string) (string, error)
	DeleteField(field_id string) (string, error)
	DeleteImage(image_id string) (string, error)
}