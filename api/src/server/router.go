package server

import (
	"net/http"

	"github.com/api/src/config"
	"github.com/api/src/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: c.GetStringSlice("server.cors"),
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	version := router.Group("/" + c.GetString("server.version"))

	healthController := controllers.NewHealthController()
	version.GET("/health", healthController.Index)

	return router, nil
}
