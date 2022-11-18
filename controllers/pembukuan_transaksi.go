package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/models"
)

func PenutupanPembukuan(c echo.Context) error {
	tanggal := c.FormValue("tanggal")

	result, err := models.Penutupan_Pembukuan(tanggal)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
