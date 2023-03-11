package postgres

import (
	"database/sql"

	"github.com/sadbard/StorageBot/internal/storage/models"
)

type DataBase struct {
	DB *sql.DB
}

func NewDataBase(db *sql.DB) *DataBase {
	return &DataBase{
		DB: db,
	}
}

func (r *DataBase) All(userId int64) ([]models.Record, error) {
	rows, err := r.DB.Query(
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

func (r *DataBase) Add(userId int64, label, login, password string) error {
	recCount, err := r.Count(userId)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(
		`
		INSERT INTO records VALUES 
		($1, $2, $3, $4, $5);
		`, userId, recCount+1, label, login, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *DataBase) Count(userId int64) (count int, err error) {
	rows, err := r.DB.Query(
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
