package postgres_test

import (
	"stad_projekt/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)


func Test_CreateRandomCountry(t *testing.T) {
	var data  =models.CountryToDB{
		Id:       uuid.NewString(),
		Name:     uuid.NewString(),
		Location: uuid.NewString(),
		CreateAt: uuid.NewString(),
	}
	_,err :=strg.Admin().Country(data) 
	require.NoError(t,err)
}