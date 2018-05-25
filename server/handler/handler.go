package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/gotify"
	"net/http"
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

func PlaylistHandler(c echo.Context) error {
	return nil
}