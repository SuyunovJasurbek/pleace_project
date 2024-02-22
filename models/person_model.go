package models

type SignInPersonModel struct {
	Lastname  string `json:"lastname" binding:"required" `
	Firstname string `json:"firstname" binding:"required" `
	Phone     string `json:"phone" binding:"required" `
	Email     string `json:"email" binding:"required" `
	Password  string `json:"password" binding:"required" `
}

type PersonCountry struct {
	PersonId   string `json:"person_id" `
	Region     string `json:"region_name"  db:"region_name" `
	District   string `json:"district_name"  db:"district_name" `
	Village    string `json:"village_name"  db:"village_name" `
	DeviceId   string `json:"device_id"  db:"device_id" `
	PleaceName string `json:"pleace_name"  db:"pleace_name" `
}

type PersonPleaceModel struct {
	DeviceId   string `json:"device_id"  db:"device_id" `
	PleaceName string `json:"pleace_name"  db:"pleace_name" `

}

type PersonSignInModel struct {
	FullName string `json:"fullname" `
	Phone    string `json:"phone" `
	Devices []PersonPleaceModel
}

type Password struct {
	Password string `json:"password" binding:"required" `
}