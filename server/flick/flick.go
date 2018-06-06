package flick

import (
	"github.com/gericass/feel-music-mock/server/dao"
	"database/sql"
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

func (l *Location) GetFlickMusic(db *sql.DB) ([]*Track, error) {
	loc, err := dao.GetLocationsByCurrentPosition(db, l.Lat, l.Lng)
	if err != nil {
		return nil, err
	}
	var trackIDs = []int64{}
	for _, v := range loc {
		tID, err := dao.GetTracksLocationsByLocationID(db, v.ID)
		if err != nil {
			return nil, err
		}
		for _, v := range tID {
			trackIDs = append(trackIDs, v.TrackID)
		}
	}
	var tracks = []*Track{}
	for _, v := range trackIDs {
		t := &dao.Track{ID: v}
		daoTrack, err := t.GetTrackByID(db)
		if err != nil {
			return nil, err
		}
		track := &Track{SpotifyID: daoTrack.SpotifyID}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (track *Track) PostFlickMusic(db *sql.DB) error {
	var lastTrackID int64
	var lastLocationID int64
	tr := &dao.Track{SpotifyID: track.SpotifyID}
	err := tr.PutTrack(db)
	if err != nil {
		return err
	}
	lastTrackID, err = dao.GetLastTrackID(db)
	if err != nil {
		return err
	}
	l := &dao.Locations{Lat: track.Location.Lat, Lng: track.Location.Lng}
	err = l.PutLocations(db)
	if err != nil {
		return err
	}
	lastLocationID, err = dao.GetLastLocationID(db)
	if err != nil {
		return err
	}
	expired := time.Now().Add(constants.FlickTrackExpiredDate)
	tl := &dao.TracksLocations{TrackID: lastTrackID, LocationID: lastLocationID, ExpiredAt: expired}
	err = tl.PutTracksLocations(db)
	if err != nil {
		return err
	}
	return nil
}
