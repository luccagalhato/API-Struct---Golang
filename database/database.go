package database

import (
	"database/sql"
	"fmt"
)

//QueryRow ...
func (s *SQLStr) QueryRow(format string, parameters ...interface{}) (*sql.Rows, error) {
	q := fmt.Sprintf(format, parameters...)
	// fmt.Println(q)
	rows, err := s.db.Query(q)
	if err == nil {
		return rows, nil
	}
	if err := s.db.Ping(); err != nil {
		return nil, err
	}
	return s.db.Query(q)
}
