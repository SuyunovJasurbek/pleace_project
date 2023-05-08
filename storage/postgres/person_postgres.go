package postgres

import (
	"github.com/jmoiron/sqlx"
)

type personRepository struct {
	db *sqlx.DB
}


func NewPersonRepository(db *sqlx.DB) *personRepository {
	return &personRepository{
		db: db,
	}
}
