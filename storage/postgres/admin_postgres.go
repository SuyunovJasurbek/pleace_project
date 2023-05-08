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
	usersTable = "users"
)

// SignIn implements storage.AdminI
func (d *adminRepository) SignIn(login string) (models.SignInDBResponse, error) {
	var resp models.SignInDBResponse
	query := `SELECT id , name , password_hash FROM ` + usersTable + ` WHERE login = $1`
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

	if resp.ID==""{
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
