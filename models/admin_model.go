package models

type SignInModel struct {
	Login    string `json:"login" binding:"required" `
	Password string `json:"password" binding:"required" `
}
type SignInDBResponse struct {
	ID string `json:"id"  db:"id" `
	Name string `json:"name"  db:"name" `
	PasswordHash string `json:"password_hash"  db:"password_hash" `
}

type SignInHandlerResponse struct {
	ID string `json:"id"  db:"id" `
	Token string `json:"token"  db:"token" `
}

type CreateStadiumModel struct {
	Phone string `json:"phone" binding:"required" `
	Location string `json:"location" binding:"required" `
	Size string `json:"size" binding:"required" `
	Price string `json:"price" binding:"required" `
	Description string `json:"description" binding:"required" `
}