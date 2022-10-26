package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/models"
	"strconv"
)

func InputInventory(c echo.Context) error {
	nama_barang := c.FormValue("nama_barang")
	jumlah_barang := c.FormValue("jumlah_barang")
	harga_barang := c.FormValue("harga_barang")

	jb, _ := strconv.Atoi(jumlah_barang)

	hb, _ := strconv.Atoi(harga_barang)

	result, err := models.Input_Inventory(nama_barang, jb, hb)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadInventory(c echo.Context) error {
	result, err := models.Read_Stock()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func Check_Nama_Inventory(c echo.Context) error {
	kode_inventory := c.FormValue("kode_inventory")
	nama_barang := c.FormValue("nama_barang")

	result, err := models.Check_Nama_Stcok(kode_inventory, nama_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
