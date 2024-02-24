package models

//___________________________ShouldBindJSON________________________________________
type SignInModel struct {
	Login    string `json:"login" binding:"required" `
	Password string `json:"password" binding:"required" `
}


type SignInHandlerResponse struct {
	Token string `json:"token"  db:"token" `
}

type AccsesUser struct {	
	Fullname string `json:"fullname"  db:"fullname" `
	PersonId string `json:"id"  db:"id" `
	Phone   string `json:"phone"  db:"phone" `
}

type Place struct {
	PlaceName string `json:"pleace_name"  db:"pleace_name" `
	PlaceId   string `json:"id"  db:"id" `
}

type PasswordResponse struct {
	Id string `json:"id"  db:"id" `
}