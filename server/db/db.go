package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mysql@tcp(db:3306)/battari?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
