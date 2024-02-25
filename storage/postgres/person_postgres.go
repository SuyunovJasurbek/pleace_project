package postgres

import (
	"fmt"
	"stad_projekt/models"
	"time"

	"github.com/google/uuid"
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

// CreatePerson implements storage.PersonI.
func (p *personRepository) SignUpPerson(data models.SignInPersonModel) error {
	fmt.Println(data)
	query := fmt.Sprintf(`INSERT INTO %s (id, fullname , phone , email , status , parol , created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`, persons_table)
	_, err := p.db.Exec(query, uuid.NewString(), data.Firstname+" "+data.Lastname, data.Phone, data.Email, "0", data.Password, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return err
	}
	return nil
}

// SignInPerson implements storage.PersonI.
func (p *personRepository) SignInPerson(parol string) (string, error) {
	var id string
	query := fmt.Sprintf(`SELECT id FROM %s WHERE parol = $1`, persons_table)
	err := p.db.Get(&id, query, parol)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return id, nil
}

// GetHumidity implements storage.PersonI.
func (p *personRepository) GetHumidity(device_id string) ([]models.GetHumidity, error) {
	var humidity []models.GetHumidity
	query := fmt.Sprintf(`SELECT result_humidity, created_at FROM %s WHERE device_id = $1`, "weather")
	err := p.db.Select(&humidity, query, device_id)
	if err != nil {
		fmt.Println(err)
		return []models.GetHumidity{}, err
	}
	return humidity, nil
}

// GetTemperature implements storage.PersonI.
func (p *personRepository) GetTemperature(device_id string) ([]models.GetTemperature, error) {
	var temperature []models.GetTemperature
	query := fmt.Sprintf(`SELECT result_sun, created_at FROM %s WHERE device_id = $1`, "weather")
	err := p.db.Select(&temperature, query, device_id)
	if err != nil {
		return []models.GetTemperature{}, err
	}
	return temperature, nil
}

// GetLight implements storage.PersonI.
func (p *personRepository) GetLight(device_id string) ([]models.GetLight, error) {
	var light []models.GetLight
	query := fmt.Sprintf(`SELECT result_wind, created_at FROM %s WHERE device_id = $1`, "weather")
	err := p.db.Select(&light, query, device_id)
	if err != nil {
		return []models.GetLight{}, err
	}
	return light, nil
}

// GetHome implements storage.PersonI.
func (p *personRepository) GetHome(id string ) (models.PersonSignInModel, error) {
	var person models.PersonSignInModel
	query1 := `select fullname, phone from persons where id = $1`
	err := p.db.Get(&person, query1, id)
	if err != nil {
		fmt.Println(err)
		return models.PersonSignInModel{}, err
	}

	query2 := `select device_id , pleace_name from country where person_id = $1`
	err = p.db.Select(&person.Devices, query2, id)
	if err != nil {
		fmt.Println(err,12121212)
		return models.PersonSignInModel{}, err
	}
	return person, nil
}

// getDayDate implements storage.PersonI.
func (p *personRepository) GetDayDate(device_id string) ([]models.DayDate, error) {
	var dayDate []models.DayDate
	query := fmt.Sprintf(`SELECT result_humidity, result_sun, result_wind , created_at  FROM %s WHERE device_id = $1`, "weather")
	err := p.db.Select(&dayDate, query, device_id)
	if err != nil {
		return []models.DayDate{}, err
	}
	return dayDate, nil
}