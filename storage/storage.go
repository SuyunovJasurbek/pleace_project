package storage

import "stad_projekt/models"

type StorageI interface {
	Person() PersonI
	Admin() AdminI
	CloseDb()
}

type PersonI interface{
}

type AdminI interface {
	SignIn(models.SignInModel)(  string ,error)
	Auth(token string) (bool)
	UserList(models.UserList) ([]models.UserList, error)
	Country(models.CountryToDB)( string, error)
	Feild(models.FeildToDB) (string, error)
	Picture(models.PictureToDB) (string, error)
	GetCountry() ([]models.GetCountry, error)
	GetField() ([]models.GetField, error)
	CreateData(models.AparatDataToDB) (string, error)
	GetData(models.GetFeildId) ([]models.AparatDataToDB, error)
	GetPicture(models.GetFeildId) ([]models.GetPicture, error)
}