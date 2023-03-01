package storage

import (
	"database/sql"
	"fmt"
)

type RecordInterface interface {
	All() ([]Record, error)
}

type Record struct {
	UserID int
	Name   string
	Login  string
	Passw  string
}

type RecordDB struct {
	DB *sql.DB
}

func (r *RecordDB) All() ([]Record, error) {
	rows, err := r.DB.Query("SELECT * FROM records")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var records []Record

	for rows.Next() {
		var rec Record
		err := rows.Scan(&rec.UserID, &rec.Name, &rec.Login, &rec.Passw)
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
	return fmt.Sprintf("%d   %s   %s   %s\n", r.UserID, r.Name, r.Login, r.Passw)
}
