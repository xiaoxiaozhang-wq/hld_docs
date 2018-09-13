package main

import (
	"flag"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	listen string
)

func init() {
	flag.StringVar(&listen, "l", ":1323", "-l addr")
}

func main() {
	flag.Parse()
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	// e.Use(middleware.Static("/_book"))
	// the file server for rice. "app" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("_book").HTTPBox())
	// serves the index.html from rice
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))

	e.Logger.Fatal(e.Start(listen))
}
