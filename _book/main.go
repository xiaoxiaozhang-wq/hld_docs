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
	dir    string
)

func init() {
	flag.StringVar(&listen, "l", ":1323", "-l addr")
	flag.StringVar(&dir, "d", "_book", "-d dir")

}

func main() {
	flag.Parse()
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	// e.Use(middleware.Static("/_book"))
	// the file server for rice. "app" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("_book").HTTPBox())

	// riceCfg := rice.Config{LocateOrder: []rice.LocateMethod{rice.LocateWorkingDirectory, rice.LocateAppended, rice.LocateEmbedded}}
	// // the file server for rice. "app" is the folder where the files come from.
	// assetHandler := http.FileServer(riceCfg.MustFindBox(dir).HTTPBox())
	// serves the index.html from rice
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", assetHandler)))

	e.Logger.Fatal(e.Start(listen))
}
