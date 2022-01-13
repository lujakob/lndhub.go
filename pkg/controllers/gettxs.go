package controllers

import (
	"github.com/bumi/lndhub.go/pkg/database/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetTXSController : GetTXSController struct
type GetTXSController struct{}

// GetTXS : Get TXS Controller
func (GetTXSController) GetTXS(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.Invoice{})
}
