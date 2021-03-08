package controllers

import (
	"net/http"

	"github.com/cecep31/goecho/models"
	"github.com/labstack/echo/v4"
)

func GetArtikel(c echo.Context) error {
	var js models.Artikel
	j := &js
	return c.JSON(http.StatusOK, j)
}
