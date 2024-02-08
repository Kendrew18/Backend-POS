package stock_masuk

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/stock_masuk"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputStockMasuk(c echo.Context) error {
	var Request request.Input_Stock_Masuk_Request
	var Request_barang request.Input_Barang_Stock_Masuk_Request

	Request.Kode_supplier = c.FormValue("kode_supplier")
	Request.Kode_user = c.FormValue("kode_user")
	Request.Nama_penanggung_jawab = c.FormValue("nama_penanggung_jawab")
	Request.Tanggal_masuk = c.FormValue("tanggal_masuk")

	Request_barang.Kode_stock = c.FormValue("kode_stock")
	Request_barang.Jumlah = c.FormValue("jumlah")
	Request_barang.Harga = c.FormValue("harga")

	result, err := stock_masuk.Input_Stock_Masuk(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadStockMasuk(c echo.Context) error {
	var Request request.Read_Stock_Masuk_Request
	var Request_filter request.Read_Stock_Masuk_Filter_Request

	Request.Kode_user = c.FormValue("kode_user")

	Request_filter.Kode_supplier = c.FormValue("kode_supplier")
	Request_filter.Tanggal_awal = c.FormValue("tanggal_awal")
	Request_filter.Tanggal_akhir = c.FormValue("tanggal_akhir")

	result, err := stock_masuk.Read_Stock_Masuk(Request, Request_filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

// func UpdateBarangStockMasuk(c echo.Context) error {
// 	var Request request.Update_Stock_Masuk_Request
// 	var Request_kode request.Update_Kode_Barang_Stock_Masuk_Request

// 	Request.Kode_stock = c.FormValue("kode_stock")
// 	Request.Jumlah, _ = strconv.ParseFloat(c.FormValue("jumlah"), 64)
// 	Request.Harga, _ = strconv.ParseInt(c.FormValue("harga"), 10, 64)

// 	Request_kode.Kode_barang_stock_masuk = c.FormValue("kode_barang_stock_masuk")

// 	result, err := stock_masuk.Update_Barang_Stock_Masuk(Request, Request_kode)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(result.Status, result)
// }

// func DeleteBarangStockMasuk(c echo.Context) error {
// 	var Request request.Update_Kode_Barang_Stock_Masuk_Request

// 	Request.Kode_barang_stock_masuk = c.FormValue("kode_barang_stock_masuk")

// 	result, err := stock_masuk.Delete_Barang_Stock_Masuk(Request)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(result.Status, result)
// }

// func DeleteBarangStockMasuk(c echo.Context) error {
// 	var Request request.Update_Kode_Barang_Stock_Masuk_Request

// 	Request.Kode_barang_stock_masuk = c.FormValue("kode_barang_stock_masuk")

// 	result, err := stock_masuk.Update_Status_Stock_Masuk(Request)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(result.Status, result)
// }

// func Read_Detail_Stock_Masuk(c echo.Context) error {
// 	id_stock_masuk := c.FormValue("id_stock_masuk")

// 	result, err := models.Read_Detail_Stock_Masuk(id_stock_masuk)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
