package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/gotify"
	"net/http"
	"github.com/gericass/feel-music-mock/server/handler/model"
	"github.com/gericass/feel-music-mock/server/flick"
)

var Auth gotify.OAuth
var Token gotify.Gotify

// Handler : Controller for https://localhost:3000/
func AuthHandler(c echo.Context) error {
	url := Auth.AuthURL() // Get the Redirect URL for authenticate
	return c.Redirect(301, url)
}

// CallbackHandler : Controller for https://localhost:3000/callback/
func CallbackHandler(c echo.Context) error {
	t, err := Auth.GetToken(c.Request()) // Get the token for using Spotify API
	if err != nil {
		return err
	}
	Token = t
	return c.String(http.StatusOK, "Authentication success")
}

func PlotTracksMapHandler(c echo.Context) error {
	return nil
}

func GetTracksFromMapHandler(c echo.Context) error {
	return nil
}

func PostFlickTrackHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	u := new(model.Flick)
	err := c.Bind(u)
	if err != nil {
		return err
	}
	track := &flick.Track{SpotifyID: u.SpotifyID, Location: flick.Location{Lat: u.Location.Lat, Lng: u.Location.Lng}}
	err = track.PostFlickMusic(cc.DB)
	if err != nil {
		return err
	}
	return c.String(200, "Track received")
}

func GetFlickTrackHandler(c echo.Context) error {
	return nil
}
