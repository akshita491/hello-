package main

import "github.com/labstack/echo/v4"


func main() {

	api := echo.New()
	api.GET("/", func(c echo.Context) error {
		return c.String(200, "hello world")
	})
	api.Start(":9090")
}