package main

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/ohwoo-kwon/learn-go/scrapper"
)

func handleHome (c echo.Context) error {
	return c.File("home.html")
}

func handleScrape (c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	return nil
}

func main () {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}