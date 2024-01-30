package controllers

import (
	"net/http"
	"project-1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputStock(c echo.Context) error {
	nama_barang := c.FormValue("nama_barang")
	jumlah_barang := c.FormValue("jumlah_barang")
	harga_barang := c.FormValue("harga_barang")
	satuan_barang := c.FormValue("satuan_barang")

	jb, _ := strconv.ParseFloat(jumlah_barang, 64)

	hb, _ := strconv.Atoi(harga_barang)

	result, err := models.Input_Inventory(nama_barang, jb, hb, satuan_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadStock(c echo.Context) error {
	result, err := models.Read_Stock()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStock(c echo.Context) error {
	kode_stock := c.FormValue("kode_stock")
	nama_barang := c.FormValue("nama_barang")
	jumlah_barang := c.FormValue("jumlah_barang")
	harga_barang := c.FormValue("harga_barang")
	satuan_barang := c.FormValue("satuan_barang")

	jb, _ := strconv.ParseFloat(jumlah_barang, 64)

	hb, _ := strconv.Atoi(harga_barang)

	result, err := models.Update_Stock(kode_stock, nama_barang, jb, hb, satuan_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CheckNamaStock(c echo.Context) error {
	kode_stock := c.FormValue("kode_stock")
	nama_barang := c.FormValue("nama_barang")

	result, err := models.Check_Nama_Stock(kode_stock, nama_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
