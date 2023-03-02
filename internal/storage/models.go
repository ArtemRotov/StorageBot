package storage

import (
	"database/sql"
	"fmt"
)

type RecordInterface interface {
	All(userID int64) ([]Record, error)
}

type Record struct {
	ID    int
	Name  string
	Login string
	Passw string
}

type RecordDB struct {
	DB *sql.DB
}

func (r *RecordDB) All(userID int64) ([]Record, error) {
	rows, err := r.DB.Query(
		`
		SELECT R.record_id, R.name, R.login, R.password 
		FROM records R
		WHERE R.user_id = $1
		`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []Record

	for rows.Next() {
		var rec Record
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

func (r *Record) String() string {
	return fmt.Sprintf("%d   %s   %s   %s\n\n", r.ID, r.Name, r.Login, r.Passw)
}
