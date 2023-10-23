package models


type AparatData struct {
	FeildId string `json:"feild_id" binding:"required" `
	ResultHumidity string `json:"result_humidity" binding:"required" `
	ResultTemperature string `json:"result_temperature" binding:"required" `
	ResultLight string `json:"result_light" binding:"required" `
}

type AparatDataToDB struct {
	Id string
	FeildId string
	ResultHumidity string
	ResultTemperature string
	ResultLight string
	CreateAt string
}

type GetAllData struct {
List []AparatDataToDB  `json:"list" db:"list"`
}

type GetFeildId struct {
	FeildId string `json:"feild_id" binding:"required" `
}

type GetCountryId struct {
	CountryId string `json:"country_id" binding:"required" `
}

type GetFeildIdAll struct {
	List []GetFeildIdToList `json:"list" db:"list"`
}

type GetFeildIdToList struct {
	FeildId string `json:"feild_id" binding:"required" `
	Name string `json:"name" binding:"required" `
}
