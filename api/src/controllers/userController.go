package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HealthController controller for health request
type HealthController struct{}

// NewHealthController is constructer for HealthController
func NewHealthController() *HealthController {
	return new(HealthController)
}

// Index is index route for health
func (hc *HealthController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"OK",
	))
}
