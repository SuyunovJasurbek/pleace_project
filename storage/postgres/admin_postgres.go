package postgres

import (
	"fmt"
	"stad_projekt/models"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type adminRepository struct {
	db *sqlx.DB
}



var (
	users_table   = "users"
	country_table = "country"
	data_table    = "weather"
	persons_table = "persons"
)

// Auth implements storage.AdminI.
func (d *adminRepository) Auth(token string) bool {
	query := fmt.Sprintf(`SELECT token FROM %s WHERE token=$1`, users_table)
	var token_db string
	if err := d.db.DB.QueryRow(query, token).Scan(
		&token_db,
	); err != nil {
		fmt.Println("----------------------")
		fmt.Println(err)
		return false
	}
	if token_db == token {
		return true
	}
	return false
}
// GetData implements storage.AdminI.
func (d *adminRepository) GetData(dat string) ([]models.AparatDataToDB, error) {
	var cnt []models.AparatDataToDB

	query := fmt.Sprintf(`SELECT id, device_id, result_humidity, result_sun, result_wind, created_at FROM %s WHERE device_id=$1`, data_table)
	rows, err := d.db.Query(query, dat)
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.AparatDataToDB
		if err := rows.Scan(
			&c.Id,
			&c.FeildId,
			&c.ResultHumidity,
			&c.ResultTemperature,
			&c.ResultLight,
			&c.CreateAt,
		); err != nil {
			fmt.Println("____________ichida_________")
			fmt.Println(err)
			return nil, err
		}

		cnt = append(cnt, c)
	}

	return cnt, nil
}
// CreateData implements storage.AdminI.
func (d *adminRepository) CreateData(enty models.AparatDataToDB) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, device_id, result_humidity, result_sun, result_wind, created_at ) VALUES ($1, $2, $3, $4, $5, $6)`, data_table)

	if _, err := d.db.Exec(query, enty.Id, enty.FeildId, enty.ResultHumidity, enty.ResultTemperature, enty.ResultLight, enty.CreateAt); err != nil {
		fmt.Println("_________________________")
		fmt.Println(err)
		return "", err
	}

	return enty.Id, nil
}
// SignIn implements storage.AdminI.
func (d *adminRepository) SignIn(ent models.SignInModel) (string, error) {
	var token string

	qurty := fmt.Sprintf(`SELECT token  FROM %s  WHERE login =$1 and  password_hash=$2`, users_table)
	if err := d.db.DB.QueryRow(qurty, ent.Login, ent.Password).Scan(
		&token,
	); err != nil {
		fmt.Println("----------------------")
		fmt.Println(err)
		return "", err

	}

	token_new := uuid.NewString()

	query2 := fmt.Sprintf(`UPDATE %s SET token=$1 WHERE login=$2`, users_table)
	if _, err := d.db.Exec(query2, token_new, ent.Login); err != nil {
		fmt.Println("_________________________")
		fmt.Println(err)
		return "", err
	}

	return token_new, nil
}
// InactiveUsers implements storage.AdminI.
func (d *adminRepository) GetInactiveUsers() ([]models.AccsesUser, error) {
	query := `SELECT  fullname, id FROM persons where status =  '0' ;`
	var cnt []models.AccsesUser
	rows, err := d.db.Query(query)
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.AccsesUser
		if err := rows.Scan(
			&c.Fullname,
			&c.PersonId,
		); err != nil {
			fmt.Println("____________ichida_________")
			fmt.Println(err)
			return nil, err
		}

		cnt = append(cnt, c)
	}

	return cnt, nil
}
// ActiveUsers implements storage.AdminI.
func (d *adminRepository) GetActiveUsers() ([]models.AccsesUser, error) {
	query := `SELECT  fullname, id FROM persons where status =  '1' ;`
	var cnt []models.AccsesUser
	rows, err := d.db.Query(query)
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.AccsesUser
		if err := rows.Scan(
			&c.Fullname,
			&c.PersonId,
		); err != nil {
			fmt.Println("____________ichida_________")
			fmt.Println(err)
			return nil, err
		}

		cnt = append(cnt, c)
	}

	return cnt, nil
}
// CreatePersonCountry implements storage.AdminI.
func (d *adminRepository) CreatePersonCountry(data models.PersonCountry) error {
	fmt.Println(data)
	query := fmt.Sprintf(`INSERT INTO %s (id, person_id ,  region_name , district_name ,  village_name , device_id , status , pleace_name , created_at ) VALUES ($1, $2, $3, $4, $5, $6 , $7,$8,$9)`, country_table)
	_, err := d.db.Exec(query, uuid.NewString(),data.PersonId, data.Region, data.District, data.Village, data.DeviceId, "1", data.PleaceName, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	update_query :=`UPDATE persons SET status = '1' WHERE id = $1`
	_, err = d.db.Exec(update_query, data.PersonId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
// GetActivePleaces implements storage.AdminI.
func (d *adminRepository) GetActivePleaces(person_id string) ([]models.Place, error) {
	fmt.Println(person_id)
	query := `SELECT  id, pleace_name FROM country where status =  '1' and person_id =$1 ;`
	var cnt []models.Place
	rows, err := d.db.Query(query, person_id)
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Place
		if err := rows.Scan(
			&c.PlaceId,
			&c.PlaceName,
		); err != nil {
			fmt.Println("____________ichida_________")
			fmt.Println(err)
			return nil, err
		}

		cnt = append(cnt, c)
	}

	return cnt, nil
}
// UpdatePleace implements storage.AdminI.
func (d *adminRepository) UpdatePleace(pleace_id string) error {
	query := `UPDATE country SET status = '0' WHERE id = $1`
	_, err := d.db.Exec(query, pleace_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
// UpdatePerson implements storage.AdminI.
func (d *adminRepository) UpdatePerson(person_id string) error {
	query := `UPDATE persons SET status = '0' WHERE id = $1`
	_, err := d.db.Exec(query, person_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewAdminRepository(db *sqlx.DB) *adminRepository {
	return &adminRepository{
		db: db,
	}
}