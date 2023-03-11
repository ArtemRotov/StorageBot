package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sadbard/StorageBot/internal/storage/models"
)

type DataBase struct {
	DB         *sql.DB
	dbUser     string
	dbPassword string
	dbName     string
	dbSSLM     string
}

func openDB(user, pass, name, ssl string) *sql.DB {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		user, pass, name, ssl)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	return db
}

func NewDataBase() *DataBase {
	db := &DataBase{
		DB:         nil,
		dbUser:     os.Getenv("DB_USER"),
		dbPassword: os.Getenv("DB_PASSW"),
		dbName:     os.Getenv("DB_NAME"),
		dbSSLM:     os.Getenv("BS_SSLM"),
	}

	db.DB = openDB(db.dbUser, db.dbPassword, db.dbName, db.dbSSLM)

	return db
}

func (d *DataBase) Close() {
	d.DB.Close()
}

func (d *DataBase) All(userId int64) ([]models.Record, error) {
	rows, err := d.DB.Query(
		`
		SELECT R.rec_id, R.rec_label, R.login, R.password
		FROM records R
		WHERE R.user_id = $1
		`, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []models.Record

	for rows.Next() {
		var rec models.Record
		err := rows.Scan(&rec.ID, &rec.Name, &rec.Login, &rec.Passw)
		if err != nil {
			return nil, err
		}

		records = append(records, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (d *DataBase) Add(userId int64, label, login, password string) error {
	recCount, err := d.Count(userId)
	if err != nil {
		return err
	}

	_, err = d.DB.Exec(
		`
		INSERT INTO records VALUES 
		($1, $2, $3, $4, $5);
		`, userId, recCount+1, label, login, password)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataBase) Count(userId int64) (count int, err error) {
	rows, err := d.DB.Query(
		`
		SELECT COUNT(R.rec_id) FROM records R
		WHERE R.user_id = $1
		`, userId)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return count, err
}
