package models

import (
	"fmt"
)

type Record struct {
	ID    int
	Name  string
	Login string
	Passw string
}

func (r *Record) String() string {
	return fmt.Sprintf("%d   %s   %s   %s\n\n", r.ID, r.Name, r.Login, r.Passw)
}
