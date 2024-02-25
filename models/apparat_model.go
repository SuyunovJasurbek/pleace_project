package models


type AparatData struct {
	FeildId string `json:"device_id" binding:"required" `
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

type GetHumidity struct {
	Humidity string `json:"humidity" db:"result_humidity" `
	Date string `json:"date" db:"created_at" `
}

type GetTemperature struct {
	Temperature string `json:"temperature" db:"result_sun" `
	Date string `json:"date" db:"created_at" `
}

type GetLight struct {
	Light string `json:"light" db:"result_wind" `
	Date string `json:"date" db:"created_at" `
}

type HomeDataStatic struct {
	Humidity []GetHumidity `json:"humidity" db:"humidity" `
	Tempreature []GetTemperature `json:"temperature" db:"temperature" `
	Light []GetLight `json:"light" db:"light" `
}

type DayDate struct {
	Day string `json:"day" binding:"required" `
	Temperature string `json:"temperature" db:"result_sun" `
	Light string `json:"light" db:"result_wind" `
	Humidity string `json:"humidity" db:"result_humidity" `
	Date string `json:"date" db:"created_at" `
}

type GetDeviceIDData struct {
	Humidity []GetHumidity `json:"humidity" db:"humidity" `
	Tempreature []GetTemperature `json:"temperature" db:"temperature" `
	Light []GetLight `json:"light" db:"light" `
	DayDate []DayDate `json:"day_date" db:"day_date" `
}