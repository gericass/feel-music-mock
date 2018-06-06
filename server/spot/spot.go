package spot

import (
	"database/sql"
	"github.com/gericass/feel-music-mock/server/dao"
	"time"
	"github.com/gericass/feel-music-mock/server/constants"
)

type Location struct {
	Lat float64
	Lng float64
}

type Track struct {
	SpotifyID string
	Location  Location
}

func (track *Track) PostMapMusic(db *sql.DB) error {
	var lastTrackID int64
	var lastMapID int64
	tr := &dao.Track{SpotifyID: track.SpotifyID}
	err := tr.PutTrack(db)
	if err != nil {
		return err
	}
	lastTrackID, err = dao.GetLastTrackID(db)
	if err != nil {
		return err
	}
	lc := &dao.Map{Lat: track.Location.Lat, Lng: track.Location.Lng}
	err = lc.PutMap(db)
	if err != nil {
		return err
	}
	lastMapID, err = dao.GetLastMapID(db)
	if err != nil {
		return err
	}
	dt := time.Now().Add(constants.MapTrackExpiredDate)
	mt := &dao.MapsTracks{TrackID: lastTrackID, MapID: lastMapID, ExpiredAt: dt}
	err = mt.PutMapsTracks(db)
	if err != nil {
		return err
	}
	return nil
}
