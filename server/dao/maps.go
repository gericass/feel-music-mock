package dao

import (
	"database/sql"
	"time"
)

type Map struct {
	ID        int64
	Lat       float64
	Lng       float64
	CreatedAt time.Time
}

func (m *Map) PutMap(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "INSERT INTO `maps` (`lat`,`lng`) values(?,?)"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(m.Lat, m.Lng)
		if err != nil {
			return err
		}
		return nil
	})
}

func GetMaps(db *sql.DB) ([]*Map, error) {
	maps := make([]*Map, 0)
	rows, err := db.Query("SELECT `id`, `lat`, `lng` FROM `maps`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		location := &Map{}
		if err := rows.Scan(&location.ID, &location.Lat, &location.Lng, &location.CreatedAt); err != nil {
			return nil, err
		}
		maps = append(maps, location)
	}
	return maps, nil
}

func GetLastMapID(db *sql.DB) (int64, error) {
	var id int64
	rows, err := db.Query("SELECT max(id) FROM `maps`")
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}
	return id, nil
}
