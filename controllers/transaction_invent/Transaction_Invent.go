package transaction_invent

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/transaction_inventory"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputTransactionInventory(c echo.Context) error {
	var Request request.Input_Transaksi_Inventory_Request
	var Request_barang request.Input_Barang_Transaksi_Inventory_Request

	Request.Tanggal = c.FormValue("tanggal")
	Request.Kode_transaksi = c.FormValue("kode_transaksi")
	Request.Kode_jenis_pembayaran = c.FormValue("kode_jenis_pembayaran")
	Request.Harga_ongkos_kirim, _ = strconv.ParseInt((c.FormValue("harga_ongkos_kirim")), 10, 64)
	Request.Ppn, _ = strconv.ParseFloat(c.FormValue("ppn"), 64)
	Request.Kode_supplier = c.FormValue("kode_supplier")
	Request.Kode_user = c.FormValue("kode_user")
	Request.Jenis_transaksi = c.FormValue("jenis_transaksi")

	Request_barang.Kode_inventory = c.FormValue("kode_inventory")
	Request_barang.Jumlah = c.FormValue("jumlah")
	Request_barang.Harga = c.FormValue("harga")

	result, err := transaction_inventory.Input_Transaction_Inventory(Request, Request_barang)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

// func ReadTransactionInventory(c echo.Context) error {
// 	var Request request.Read_Stock_Masuk_Request
// 	var Request_filter request.Read_Stock_Masuk_Filter_Request

// 	result, err := transaction_inventory.Read_Transaction_Inventory(Request, Request_filter)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(result.Status, result)
// }
