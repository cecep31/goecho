package main

import (
	"net/http"

	"github.com/cecep31/goecho/config"
	"github.com/cecep31/goecho/controllers"
	"github.com/cecep31/goecho/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Dbcon().AutoMigrate(&models.Artikel{})
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339},` +
			`"${method}","uri":"${uri}"` +
			`|${status},"latency":${latency}|${error}` +
			"\n",
	}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	config.Dbcon()
	// Routes
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.GET("/post", controllers.GetArtikels)
	e.POST("/post", controllers.AddArtikel)
	e.GET("/post/:id", controllers.GetArtikel)
	e.DELETE("/post/:id", controllers.DeleteArtikel)
	e.PATCH("/post/:id", controllers.UpdateArtikel)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
