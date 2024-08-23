package main

import (
	"database/sql"
)

func healthFunc(db *sql.DB) func() error {
	return func() error {
		if err := db.Ping(); err != nil {
			return err
		}
		return nil
	}
}
