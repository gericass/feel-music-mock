package dao

import (
	"time"
	"database/sql"
)

type TracksLocations struct {
	ID         int64
	TrackID    int64
	LocationID int64
	CreatedAt  time.Time
	ExpiredAt  time.Time
}

func (tl *TracksLocations) PutTracksLocations(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "INSERT INTO `tracks_locations` (`track_id`, `location_id`, `expired_at`) values(?,?,?)"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(tl.TrackID, tl.LocationID, tl.ExpiredAt)
		if err != nil {
			return err
		}
		return nil
	})
}

func GetTracksLocationsByLocationID(db *sql.DB, locationID int64) ([]*TracksLocations, error) {
	tls := make([]*TracksLocations, 0)
	stmt, err := db.Prepare("SELECT `id`, `track_id`, `location_id`, `created_at`, `expired_at` FROM `tracks_locations` WHERE `location_id` = ?")
	rows, err := stmt.Query(stmt, locationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		tl := &TracksLocations{}
		if err := rows.Scan(&tl.ID, &tl.TrackID, &tl.LocationID, &tl.CreatedAt, &tl.ExpiredAt); err != nil {
			return nil, err
		}
		tls = append(tls, tl)
	}
	return tls, nil
}
