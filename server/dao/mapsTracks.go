package dao

import (
	"time"
	"database/sql"
)

type MapsTracks struct {
	ID        int64
	TrackID   int64
	MapID     int64
	CreatedAt time.Time
	ExpiredAt time.Time
}

func (t *MapsTracks) PutMapsTracks(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "INSERT INTO `maps_tracks` (`track_id`, `map_id`, `expired_at`) values(?,?,?)"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(t.TrackID, t.MapID, t.ExpiredAt)
		if err != nil {
			return err
		}
		return nil
	})
}

func GetMapsTracksByMapID(db *sql.DB, mapId int) ([]*MapsTracks, error) {
	mapsTracks := make([]*MapsTracks, 0)
	stmt, err := db.Prepare("SELECT `id`, `track_id`, `map_id`, `created_at`, `expired_at` FROM `maps_tracks` WHERE `map_id` = ?")
	rows, err := stmt.Query(stmt, mapId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mt := &MapsTracks{}
		if err := rows.Scan(&mt.ID, &mt.TrackID, &mt.MapID, &mt.CreatedAt, &mt.ExpiredAt); err != nil {
			return nil, err
		}
		mapsTracks = append(mapsTracks, mt)
	}
	return mapsTracks, nil
}
