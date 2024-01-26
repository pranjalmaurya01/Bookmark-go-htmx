package main

import (
	"bookmark/models"
	"bookmark/templates/components"
	"bookmark/templates/pages"
	"bookmark/utilities"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Website{})

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		var websites []models.Website
		err := db.Find(&websites).Error
		if err != nil {
			fmt.Println("err", err)
		}
		component := pages.Home(websites)
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.POST("/add-website", func(c echo.Context) error {
		link := c.FormValue("link")
		fmt.Println("link", link)
		meta, e := utilities.GetMetaData(link)

		if !e {
		}
		var websites []models.Website
		err := db.Find(&websites).Error
		if err != nil {
			fmt.Println("err", err)
		}

		newWebsite := models.Website{Name: meta.Title, Description: meta.Description, Img: meta.Image, Url: link, UserId: 1}

		fmt.Println(newWebsite)
		db.Save(&newWebsite)

		component := components.Card(newWebsite)
		return component.Render(context.Background(), c.Response().Writer)
	})

	port := utilities.GetEnvVar("PORT")
	if port == "" {
		port = "3000"
	}
	e.Static("/static", "public")
	e.Logger.Fatal(e.Start(":" + port))

}
