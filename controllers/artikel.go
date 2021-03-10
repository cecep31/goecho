package controllers

import (
	"net/http"

	"github.com/cecep31/goecho/config"
	"github.com/cecep31/goecho/models"
	"github.com/labstack/echo/v4"
)

func GetArtikels(c echo.Context) error {
	var data []models.Artikel
	db := config.Dbcon()
	db.Find(&data)

	return c.JSON(http.StatusOK, data)
}

func GetArtikel(c echo.Context) error {
	var data models.Artikel
	if err := config.Dbcon().Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

type atikelinputadd struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func AddArtikel(c echo.Context) error {
	var input atikelinputadd
	if err := c.Bind(&input); err != nil {
		return err
	}
	art := models.Artikel{Title: input.Title, Content: input.Content}
	config.Dbcon().Create(&art)

	return c.JSON(http.StatusCreated, art.Title)

}
