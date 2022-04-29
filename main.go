package main

import "github.com/labstack/echo/v4"
import "math/rand"
import "net/http"

type quote struct {
	Title string
	Author string
}
var quotes []quote =[]quote {
	{"Naruto Shippuden","the OOOOOFFFIIcee"},
	// fo random input more values 
}

func main() {
	api := echo.New()
	api.GET("/quotes",getRandomQuotes)
	api.Start(":9093")
}

func getRandomQuotes(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}