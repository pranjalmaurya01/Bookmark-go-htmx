package main

import (
	"bookmark/templates/pages"
	"bookmark/utilities"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		component := pages.Home()
		return component.Render(context.Background(), c.Response().Writer)
	})

	port := utilities.GetEnvVar("PORT")
	if port == "" {
		port = "3000"
	}
	e.Static("/static", "public")
	e.Logger.Fatal(e.Start(":" + port))
}
