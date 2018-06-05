package db

import "database/sql"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/feel-music-mock?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
