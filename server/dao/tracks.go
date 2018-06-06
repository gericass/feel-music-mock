package dao

import (
	"database/sql"
	"time"
)

type Track struct {
	ID        int64
	SpotifyID string
	CreatedAt time.Time
}

func (t *Track) PutTrack(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "INSERT INTO `tracks` (`spotify_id`) values(?)"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(t.SpotifyID)
		if err != nil {
			return err
		}
		return nil
	})
}

func (t *Track) GetTrackByID(db *sql.DB) (*Track, error) {
	stmt, err := db.Prepare("SELECT `id`, `spotify_id`, `created_at` FROM `tracks` WHERE `id` = ?")
	if err != nil {
		return nil, err
	}
	track := &Track{}
	err = stmt.QueryRow(stmt, t.ID).Scan(&track.ID, &track.SpotifyID, &track.CreatedAt)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func GetLastTrackID(db *sql.DB) (int64, error) {
	var id int64
	rows, err := db.Query("SELECT max(id) FROM `tracks`")
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
