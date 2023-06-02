package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(192.168.64.2:3306)/cijou")
	if err != nil {
		return nil, err
	}
	return db, nil
}
