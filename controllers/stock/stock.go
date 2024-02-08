package stock

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/stock"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputStock(c echo.Context) error {
	var Request request.Input_Stock_Request

	Request.Nama_barang = c.FormValue("nama_barang")
	Request.Jumlah_barang, _ = strconv.ParseFloat(c.FormValue("jumlah_barang"), 64)
	Request.Harga_barang, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)
	Request.Satuan_barang = c.FormValue("satuan_barang")
	Request.Kode_user = c.FormValue("kode_user")

	result, err := stock.Input_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadStock(c echo.Context) error {
	var Request request.Read_Stock_Request

	Request.Kode_user = c.FormValue("kode_user")

	result, err := stock.Read_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateStock(c echo.Context) error {
	var Request request.Update_Stock_Request

	Request.Kode_stock = c.FormValue("kode_stock")
	Request.Nama_barang = c.FormValue("nama_barang")
	Request.Jumlah_barang, _ = strconv.ParseFloat(c.FormValue("jumlah_barang"), 64)
	Request.Harga_barang, _ = strconv.ParseInt(c.FormValue("harga_barang"), 10, 64)
	Request.Satuan_barang = c.FormValue("satuan_barang")

	result, err := stock.Update_Stock(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

// func CheckNamaStock(c echo.Context) error {
// 	kode_stock := c.FormValue("kode_stock")
// 	nama_barang := c.FormValue("nama_barang")

// 	result, err := models.Check_Nama_Stock(kode_stock, nama_barang)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
