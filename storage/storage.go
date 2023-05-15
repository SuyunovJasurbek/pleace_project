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
	SignIn(login string)(models.SignInDBResponse, error)
	CreateStadiumName(models.StadiumNameRequest)(string , error)
	CheckStadiumPicture(models.CheckStadiumPicture)(bool)
	UploadsPictures()
}
