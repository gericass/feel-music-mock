package dao

import (
	"time"
	"database/sql"
)

type Locations struct {
	ID        int64
	Lat       float64
	Lng       float64
	CreatedAt time.Time
}

func (l *Locations) PutLocations(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "INSERT INTO `locations` (`lat`, `lng`) values(?, ?)"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(l.Lat, l.Lng)
		if err != nil {
			return err
		}
		return nil
	})
}

func GetLocationsByCurrentPosition(db *sql.DB, lat float64, lng float64) ([]*Locations, error) {
	ls := make([]*Locations, 0)
	stmt, err := db.Prepare(`SELECT
  id, lat, lng, created_at
  (
    6371 * acos( -- kmの場合は6371、mileの場合は3959
      cos(radians(?))
      * cos(radians(lat))
      * cos(radians(lng) - radians(?))
      + sin(radians(?))
      * sin(radians(lat))
    )
  ) AS distance
FROM
  locations
HAVING
  distance <= 3 -- 半径3km以内
ORDER BY
  distance -- 近い順(ASC)
LIMIT 10 -- 最大10件
;`)
	rows, err := stmt.Query(stmt, lat, lng, lat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		l := &Locations{}
		if err := rows.Scan(&l.ID, &l.Lat, &l.Lng, &l.CreatedAt); err != nil {
			return nil, err
		}
		ls = append(ls, l)
	}
	return ls, nil
}


func GetLastLocationID(db *sql.DB) (int64, error) {
	var id int64
	rows, err := db.Query("SELECT max(id) FROM `locations`")
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