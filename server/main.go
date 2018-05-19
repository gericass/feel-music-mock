package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/feel-music-mock/server/handler"
)

func main() {
	e := echo.New()
	e.GET("/playlist", handler.PlaylistHandler)
}
