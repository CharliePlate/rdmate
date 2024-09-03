package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("player_api.php", respond)
	e.GET("/movie/*/*/*", movie)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func respond(c echo.Context) error {
	return nil
}

func movie(c echo.Context) error {
	return nil
}
