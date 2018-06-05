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

func GetTracks(db *sql.DB) ([]*Track, error) {
	tracks := make([]*Track, 0)
	rows, err := db.Query("SELECT `id`, `spotify_id`, `created_at` FROM `tracks`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		track := &Track{}
		if err := rows.Scan(&track.ID, &track.SpotifyID, &track.CreatedAt); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}
