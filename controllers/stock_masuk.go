package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/models"
)

func InputStockMasuk(c echo.Context) error {
	kode_supplier := c.FormValue("kode_supplier")
	kode_stock := c.FormValue("kode_stock")
	nama_supplier := c.FormValue("nama_supplier")
	jumlah_barang := c.FormValue("jumlah_barang")
	harga_barang := c.FormValue("harga_barang")

	result, err := models.Input_Stock_Masuk(kode_supplier, kode_stock, nama_supplier, jumlah_barang, harga_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadStockMasuk(c echo.Context) error {
	result, err := models.Read_Stock_Masuk()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func Read_Detail_Stock_Masuk(c echo.Context) error {
	id_stock_masuk := c.FormValue("id_stock_masuk")

	result, err := models.Read_Detail_Stock_Masuk(id_stock_masuk)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
