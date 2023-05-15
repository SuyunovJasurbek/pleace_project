package postgres

import (
	"errors"
	"log"
	"stad_projekt/models"

	"github.com/jmoiron/sqlx"
)

type adminRepository struct {
	db *sqlx.DB
}

const (
	usersTable   = "users"
	stadiumTable = "stadium"
)

// CheckStadiumPicture implements storage.AdminI
func (d *adminRepository) CheckStadiumPicture(rq models.CheckStadiumPicture) bool {
	var stadion_id, user_id string
	//1 .....
	query := ` SELECT id FROM ` + stadiumTable + ` WHERE id = $1 `
	err := d.db.QueryRow(query, rq.StadiumId).Scan(&stadion_id)
	if err != nil {
		log.Printf("Method: CheskPictures, Error: %v", err)
		return false
	}

	if stadion_id == "" {
		log.Printf("Method: CheskPictures nill stadium_id, Error: %v", err)
		return false
	}
	//2 .....
	query2 := `SELECT user_id FROM ` + stadiumTable + ` WHERE id = $1 `

	err2 := d.db.QueryRow(query2, rq.StadiumId).Scan(&user_id)

	if err2 != nil {
		log.Printf("Method: CheskPictures, Error: %v", err2)
		return false
	}

	if user_id == "" {
		log.Printf("Method: CheskPictures, user_id nil Error: %v", err2)
		return false
	}

	if user_id != rq.User_Id {
		log.Printf("Method: CheskPictures.  cleint user_id vs user_id no some ))) , Error: %v", err2)
		return false
	}

	return true
}

// UploadsPictures implements storage.AdminI
func (*adminRepository) UploadsPictures() {
	panic("unimplemented")
}

// CreateStadiumName implements storage.AdminI
func (d *adminRepository) CreateStadiumName(msg models.StadiumNameRequest) (string, error) {
	var stadium_id string
	query := `INSERT INTO ` + stadiumTable + `( id, user_id , name ) VALUES ($1,$2,$3 ) RETURNING id `
	err := d.db.QueryRow(query, msg.ID, msg.User_Id, msg.Name).Scan(&stadium_id)
	// check error
	if err != nil {
		log.Printf("Method: CreateStadiumName, Error: %v", err)
		return "", err
	}
	return stadium_id, nil
}

// SignIn implements storage.AdminI
func (d *adminRepository) SignIn(login string) (models.SignInDBResponse, error) {
	var resp models.SignInDBResponse
	query := `SELECT id , name , password_hash FROM ` + usersTable + ` WHERE login = $1 `
	// exec and scan
	err := d.db.QueryRow(
		query,
		login,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.PasswordHash,
	)

	// check error
	if err != nil {
		log.Printf("Method: SignIn, Error: %v", err)
		return models.SignInDBResponse{}, err
	}

	if resp.ID == "" {
		return models.SignInDBResponse{}, errors.New("login yoki parol xato")
	}

	// return result
	return resp, nil
}

func NewAdminRepository(db *sqlx.DB) *adminRepository {
	return &adminRepository{
		db: db,
	}
}