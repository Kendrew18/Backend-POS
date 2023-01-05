package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/models"
	"strconv"
)

func FilterTransaksi(c echo.Context) error {
	tanggal := c.FormValue("tanggal")
	tipe_status := c.FormValue("tipe_status")

	ts, _ := strconv.Atoi(tipe_status)

	result, err := models.Filter_Transaksi(tanggal, ts)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FilterStock(c echo.Context) error {
	tipe_urutan := c.FormValue("tipe_urutan")

	tu, _ := strconv.Atoi(tipe_urutan)

	result, err := models.Filter_Stock(tu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FilterStockMasuk(c echo.Context) error {
	tanggal := c.FormValue("tanggal")
	tipe_urutan := c.FormValue("tipe_urutan")
	tipe_tanggal := c.FormValue("tipe_tanggal")

	tu, _ := strconv.Atoi(tipe_urutan)
	tt, _ := strconv.Atoi(tipe_tanggal)

	result, err := models.Filter_Stock_Masuk(tanggal, tt, tu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FilterReadPembukuan(c echo.Context) error {
	tanggal := c.FormValue("tanggal")
	tanggal2 := c.FormValue("tanggal2")
	tipe_tanggal := c.FormValue("tipe_tanggal")

	tt, _ := strconv.Atoi(tipe_tanggal)

	result, err := models.Filter_Read_Pembukuan(tanggal, tanggal2, tt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
