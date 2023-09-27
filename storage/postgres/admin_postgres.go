package postgres

import (
	"fmt"
	"stad_projekt/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type adminRepository struct {
	db *sqlx.DB
}

// GetPicture implements storage.AdminI.
func (d *adminRepository) GetPicture(ent models.GetFeildId) ([]models.GetPicture, error) {
	var cnt []models.GetPicture

	query := fmt.Sprintf(`SELECT id, feild_id, path, created_at FROM %s WHERE feild_id=$1`, picture_table)
	rows, err := d.db.Query(query, ent.FeildId)
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.GetPicture
		if err := rows.Scan(
			&c.Id,
			&c.FeildId,
			&c.Url,
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

var (
	users_table   = "users"
	country_table = "country"
	feild_table   = "feild"
	picture_table = "images"
	data_table    = "weather"
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

// UserList implements storage.AdminI.
func (*adminRepository) UserList(models.UserList) ([]models.UserList, error) {
	panic("unimplemented")
}

// GetData implements storage.AdminI.
func (d *adminRepository) GetData(ent models.GetFeildId) ([]models.AparatDataToDB, error) {
	var cnt []models.AparatDataToDB

	query := fmt.Sprintf(`SELECT id, device_id, result_humidity, result_sun, result_wind, created_at FROM %s WHERE device_id=$1`, data_table)
	rows, err := d.db.Query(query, ent.FeildId)
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
	query := fmt.Sprintf(`INSERT INTO %s (id, device_id, result_humidity, result_sun, result_wind, created_at, updated_at, deleted_at ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, data_table)

	if _, err := d.db.Exec(query, enty.Id, enty.FeildId, enty.ResultHumidity, enty.ResultTemperature, enty.ResultLight, enty.CreateAt, enty.CreateAt, "-"); err != nil {
		fmt.Println("_________________________")
		fmt.Println(err)
		return "", err
	}

	return enty.Id, nil
}

// GetField implements storage.AdminI.
func (d *adminRepository) GetField() ([]models.GetField, error) {
	var cnt []models.GetField

	query := fmt.Sprintf(`SELECT id, name, country_id , created_at FROM %s`, feild_table)
	rows, err := d.db.Query(query)
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.GetField
		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.CountryId,
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

// GetCountry implements storage.AdminI.
func (d *adminRepository) GetCountry() ([]models.GetCountry, error) {
	var cnt []models.GetCountry

	query := fmt.Sprintf(`SELECT id, name, location , created_at FROM %s`, country_table)
	rows, err := d.db.Query(query)
	if err != nil {
		fmt.Println("__________tashqarida____________")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.GetCountry
		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Location,
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

// Picture implements storage.AdminI.
func (d *adminRepository) Picture(enty models.PictureToDB) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, feild_id, path, created_at, updated_at, deleted_at ) VALUES ($1, $2, $3, $4, $5 , $6)`, picture_table)

	if _, err := d.db.Exec(query, enty.Id, enty.FeildId, enty.Url, enty.CreateAt, enty.CreateAt, "-"); err != nil {
		fmt.Println("_________________________")
		fmt.Println(err)
		return "", err
	}
	return enty.Id, nil
}

// Country implements storage.AdminI.
func (d *adminRepository) Country(cnt models.CountryToDB) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, location, created_at, updated_at, deleted_at ) VALUES ($1, $2, $3, $4, $5 , $6)`, country_table)

	if _, err := d.db.Exec(query, cnt.Id, cnt.Name, cnt.Location, cnt.CreateAt, cnt.CreateAt, "-"); err != nil {
		fmt.Println("_________________________")
		fmt.Println(err)
		return "", err
	}
	return cnt.Id, nil

}

// Feild implements storage.AdminI.
func (d *adminRepository) Feild(cnt models.FeildToDB) (string, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, country_id, name, created_at, updated_at, deleted_at ) VALUES ($1, $2, $3, $4, $5 , $6)`, feild_table)

	if _, err := d.db.Exec(query, cnt.Id, cnt.CountryId, cnt.Name, cnt.CreateAt, cnt.CreateAt, "-"); err != nil {
		fmt.Println("_________________________")
		fmt.Println(err)
		return "", err
	}
	return cnt.Id, nil
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

func NewAdminRepository(db *sqlx.DB) *adminRepository {
	return &adminRepository{
		db: db,
	}
}
