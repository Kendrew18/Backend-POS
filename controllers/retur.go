package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/models"
	"strconv"
)

func InputRetur(c echo.Context) error {

	id_supplier := c.FormValue("id_supplier")
	nama_supplier := c.FormValue("id_supplier")
	kode_stock := c.FormValue("id_supplier")
	nama_barang := c.FormValue("id_supplier")
	jumlah_barang := c.FormValue("id_supplier")

	jb, _ := strconv.Atoi(jumlah_barang)

	result, err := models.Input_Retur(id_supplier, nama_supplier, kode_stock, nama_barang, jb)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadRetur(c echo.Context) error {
	result, err := models.Read_Retur()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
