package models

//___________________________ShouldBindJSON________________________________________
type SignInModel struct {
	Login    string `json:"login" binding:"required" `
	Password string `json:"password" binding:"required" `
}

type CreateStadiumModel struct {
	ID string  `json:"id" binding:"required" `
	Phone string `json:"phone" binding:"required" `
	Latitude string `json:"latitude" binding:"required" `
	Longitude string `json:"longitude" binding:"required" `
	Size string `json:"size" binding:"required" `
	Price string `json:"price" binding:"required" `
	Description string `json:"description" binding:"required" `	
	Card string  `json:"card"`	
	CardPerson string  `json:"card_person"`	
}
type CreateStadiumNameModel struct {
	Name string `json:"name" binding:"required" `
}

//________________________________HANDLER__SERVICE_______________________________________
type SignInHandlerResponse struct {
	ID string `json:"id"  db:"id" `
	Token string `json:"token"  db:"token" `
}

type CreateStadiumRequesit struct {
	Name string `json:"name" db:"name" `
	Phone string `json:"phone" db:"phone" `
	Location string `json:"location" db:"location" `
	Pictures string `json:"pictures" db:"pictures" `
	Size string `json:"size" db:"size" `
	Price string `json:"price" db:"price" `
	Description string `json:"description" db:"description" `	
	Card string  `json:"card"  db:"card"`	
	CardPerson string  `json:"card_person"  db:"card_person"`	
}

type CreateStadiumNameRequest struct {
	User_Id string `json:"user_id"  db:"user_id" `
	Name string `json:"name" db:"id" `
}

type CheckStadiumPicture struct {
	StadiumId string  `json:"stadium_id"  db:"stadium_id" `
	User_Id string  `json:"user_id"  db:"user_id" `
}

//___________________________________DB_________________________________________
type SignInDBResponse struct {
	ID string `json:"id"  db:"id" `
	Name string `json:"name"  db:"name" `
	PasswordHash string `json:"password_hash"  db:"password_hash" `
}

type SignInModelRequest struct {
	Login    string `json:"login" db:"login" `
	Password string `json:"password" db:"password" `
}
type StadiumNameRequest struct {
	ID string `json:"id"  db:"id" `
	User_Id string `json:"user_id"  db:"user_id" `
	Name string `json:"name" db:"id" `
}