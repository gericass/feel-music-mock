package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/feel-music-mock/server/handler"

	"github.com/gericass/feel-music-mock/server/db"
	"github.com/labstack/echo/middleware"
	"github.com/gericass/gotify"
)

func dbMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cnn, err := db.ConnectDB()
		if err != nil {
			return err
		}
		defer cnn.Close()
		cc := &handler.CustomContext{c, cnn}
		return h(cc)
	}
}

const clientID = ""
const clientSecret = ""
const callbackURI = "https://localhost:5000/callback/"

func main() {
	e := echo.New()
	e.Use(dbMiddleware)
	e.Pre(middleware.HTTPSRedirect())

	handler.Auth = gotify.Set(clientID, clientSecret, callbackURI)

	e.Use(middleware.Logger())

	e.GET("/playlist", handler.PlaylistHandler)
	e.GET("/auth", handler.AuthHandler)
	e.GET("/callback/", handler.CallbackHandler)

	e.Start(":5000")
}
