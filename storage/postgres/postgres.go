package postgres

import (
	"fmt"
	"log"
	"stad_projekt/config"
	"stad_projekt/storage"

	"github.com/jmoiron/sqlx"
)

type StorageI struct {
	db               *sqlx.DB
	personRepository storage.PersonI
	adminRepository  storage.AdminI
}

// Admin implements storage.StorageI
func (s *StorageI) Admin() storage.AdminI {
	if s.adminRepository == nil {
		s.adminRepository = NewAdminRepository(s.db)
	}
	return s.adminRepository
}

// CloseDb implements storage.StorageI
func (s *StorageI) CloseDb() {
	defer s.db.Close()
}

// Person implements storage.StorageI
func (s *StorageI) Person() storage.PersonI {
	if s.adminRepository == nil {
		s.personRepository = NewPersonRepository(s.db)
	}
	return s.personRepository
}

func NewPostgres(cnf config.Config) storage.StorageI {

	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.PostgresHost, cnf.PostgresPort, cnf.PostgresUser, cnf.PostgresPassword, cnf.PostgresDatabase)

	fmt.Println(psqlConnString)
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
	} else {
		log.Printf("Connected to database: %s", cnf.PostgresDatabase)
	}

	return &StorageI{
		db: db,
	}

}
