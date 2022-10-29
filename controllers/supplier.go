package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/models"
)

func InputSupplier(c echo.Context) error {
	nama_supplier := c.FormValue("nama_supplier")
	nomor_telpon := c.FormValue("nomor_telpon")
	email_supplier := c.FormValue("email_supplier")

	result, err := models.Input_Supplier(nama_supplier, nomor_telpon, email_supplier)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadSupplier(c echo.Context) error {
	result, err := models.Read_Supplier()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
