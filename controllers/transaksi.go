package controllers

// import (
// 	"github.com/labstack/echo/v4"
// 	"net/http"
// 	"project-1/models"
// 	"strconv"
// )

// func InputTransaksi(c echo.Context) error {
// 	kode_stock := c.FormValue("kode_stock")
// 	nama_barang := c.FormValue("nama_barang")
// 	jumlah_barang := c.FormValue("jumlah_barang")
// 	harga_barang := c.FormValue("harga_barang")
// 	tanggal_pelunasan := c.FormValue("tanggal_pelunasan")
// 	status_transaksi := c.FormValue("status_transaksi")
// 	sub_total_harga := c.FormValue("sub_total_harga")
// 	satuan_barang := c.FormValue("satuan_barang")

// 	sth, _ := strconv.ParseInt(sub_total_harga, 10, 64)

// 	result, err := models.Input_Transaksi(kode_stock, nama_barang, jumlah_barang, satuan_barang, harga_barang, status_transaksi, tanggal_pelunasan, sth)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

// func ReadTransaksi(c echo.Context) error {
// 	tanggal_transaksi := c.FormValue("tanggal_transaksi")

// 	result, err := models.Read_Transaksi(tanggal_transaksi)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

// func ReadDetailTransaksi(c echo.Context) error {
// 	kode_transaksi := c.FormValue("kode_transaksi")

// 	result, err := models.Read_Detail_transaksi(kode_transaksi)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

// func UpdateStatus(c echo.Context) error {
// 	kode_transaksi := c.FormValue("kode_transaksi")
// 	tanggal_pelunasan := c.FormValue("tanggal_peliunasan")

// 	result, err := models.Update_Status(kode_transaksi, tanggal_pelunasan)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

// func DateTransaksi(c echo.Context) error {
// 	tanggal := c.FormValue("tanggal")
// 	tanggal2 := c.FormValue("tanggal2")
// 	status := c.FormValue("status")

// 	st, _ := strconv.Atoi(status)

// 	result, err := models.Date_Transaksi(tanggal, tanggal2, st)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
