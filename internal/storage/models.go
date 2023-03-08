package storage

import (
	"database/sql"
	"fmt"
)

type RecordDB struct {
	DB *sql.DB
}

func NewRecordDB(db *sql.DB) *RecordDB {
	return &RecordDB{
		DB: db,
	}
}

func (r *RecordDB) All(userId int64) ([]fmt.Stringer, error) {
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

	var records []fmt.Stringer

	for rows.Next() {
		var rec Record
		err := rows.Scan(&rec.ID, &rec.Name, &rec.Login, &rec.Passw)
		if err != nil {
			return nil, err
		}

		records = append(records, &rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *RecordDB) Add(userId int64, label, login, password string) error {
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

func (r *RecordDB) Count(userId int64) (count int, err error) {
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

type Record struct {
	ID    int
	Name  string
	Login string
	Passw string
}

func (r *Record) String() string {
	return fmt.Sprintf("%d   %s   %s   %s\n\n", r.ID, r.Name, r.Login, r.Passw)
}
