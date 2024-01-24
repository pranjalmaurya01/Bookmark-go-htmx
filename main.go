package main

import (
	"bookmark/templates"
	"bookmark/utilities"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		component := templates.Hello("Pranjal")
		return component.Render(context.Background(), c.Response().Writer)
	})

	port := utilities.GetEnvVar("PORT")
	if port == "" {
		port = "3000"
	}
	e.Static("/static", "public")
	e.Logger.Fatal(e.Start(":" + port))
}
