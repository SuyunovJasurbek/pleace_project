package models

//___________________________ShouldBindJSON________________________________________
type SignInModel struct {
	Login    string `json:"login" binding:"required" `
	Password string `json:"password" binding:"required" `
}




type UserList struct {
	Name string  `json:"name" binding:"required" `
	Phone string  `json:"phone" binding:"required" `
	Login    string `json:"login" binding:"required" `
}


type SignInHandlerResponse struct {
	Token string `json:"token"  db:"token" `
}

type Country struct {
	Name string `json:"name" binding:"required" `
	Location string `json:"location" binding:"required" `
}

type CountryToDB struct {
	Id string
	Name string 
	Location string 
	CreateAt string 
}
type Feild struct {
	CountryId string `json:"country_id" binding:"required" `
	Name string `json:"name" binding:"required" `
}

type FeildToDB struct {
	Id string
	CountryId string
	Name string 
	CreateAt string 
}

type Picture struct {
	FeildId string `json:"feild_id" binding:"required" `
	Url string `json:"url" binding:"required" `
}

type PictureToDB struct {
	Id string
	FeildId string
	Url string
	CreateAt string
}

type GetCountry struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Location string `json:"location" db:"location"`
	CreateAt string `json:"create_at" db:"create_at"`
}

type GetCountryAll struct {
	List []GetCountry `json:"list" db:"list"`
}

type GetField struct {
	Id string `json:"id" db:"id"`
	CountryId string `json:"country_id" db:"country_id"`
	Name string `json:"name" db:"name"`
	CreateAt string `json:"create_at" db:"create_at"`
}

type GetFieldAll struct {
	List []GetField `json:"list" db:"list"`
}

type GetPicture struct {
	Id string `json:"id" db:"id"`
	FeildId string `json:"feild_id" db:"feild_id"`
	Url string `json:"url" db:"url"`
	CreateAt string `json:"create_at" db:"create_at"`
}

type DeleteCountryID struct {
	Id string `json:"id" binding:"required" `
}

type DeleteFieldID struct {
	Id string `json:"id" binding:"required" `
}

type UpdateCountry struct {
	Id string `json:"id" binding:"required" `
	Name string `json:"name" binding:"required" `	
	Location string `json:"location" binding:"required" `
}

type UpdateField struct {
	Id string `json:"id" binding:"required" `
	CountryId string `json:"country_id" binding:"required" `
	Name string `json:"name" binding:"required" `
}