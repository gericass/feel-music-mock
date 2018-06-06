package model

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Flick struct {
	UserName  string   `json:"userName"`
	SpotifyID string   `json:"spotifyID"`
	Location  Location `json:"location"`
}
