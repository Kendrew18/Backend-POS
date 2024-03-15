package transaction_invent

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/transaction_inventory"
	"Bakend-POS/tools/session_checking"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputTransactionInventory(c echo.Context) error {
	var Request_body request.Input_Transaksi_Body_Request
	var Request request.Input_Transaksi_Inventory_Request
	var Request_barang []request.Input_Barang_Transaksi_Inventory_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request_body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	User, condition := session_checking.Session_Checking(Request_session.Token)

	if condition {

		Request = Request_body.Transaksi_inventory
		Request_barang = Request_body.Barang_transaksi_inventory

		Request.Kode_user = User.Kode_user

		result, err = transaction_inventory.Input_Transaction_Inventory(Request, Request_barang)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

	} else {
		result.Status = http.StatusNotFound
		result.Message = "Session Invalid"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}

func ReadTransactionInventory(c echo.Context) error {

	var Request request.Read_Transaksi_Inventory_Request
	var Request_filter request.Read_Transaksi_Inventory_Filter_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")
	Request_filter.Tanggal_awal = c.Request().Header.Get("tanggal_awal")
	Request_filter.Tanggal_akhir = c.Request().Header.Get("tanggal_akhir")
	Request_filter.Nama_supplier = c.Request().Header.Get("nama_supplier")

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {

		result, err = transaction_inventory.Read_Transaction_Inventory(Request, Request_filter)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else {
		result.Status = http.StatusNotFound
		result.Message = "Session Invalid"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}

func UpdateHeaderTransactionInventory(c echo.Context) error {

	var Request_body request.Body_Update_Header_Transaksi_Inventory_Request
	var Request request.Update_Header_Transaksi_Inventory_Request
	var Request_kode request.Update_Header_Transaksi_Inventory_Kode_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request_body)

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request = Request_body.Update_header_transaksi_inventory
	Request_kode = Request_body.Update_header_transaksi_inventory_kode

	Request_kode.Kode_user = User.Kode_user

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if condition {
		result, err = transaction_inventory.Update_Header_Transaction_Inventory(Request, Request_kode)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else {
		result.Status = http.StatusNotFound
		result.Message = "Session Invalid"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}

func UpdateBarangTransactionInventory(c echo.Context) error {

	var Request_body request.Body_Update_Barang_Transaksi_Inventory
	var Request request.Update_Barang_Transaksi_Inventory_Request
	var Request_kode request.Update_Barang_Transaksi_Inventory_Kode_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request_body)

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request = Request_body.Update_barang_transaksi_inventory
	Request_kode = Request_body.Update_barang_transaksi_inventory_kode

	Request_kode.Kode_user = User.Kode_user

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if condition {

		result, err = transaction_inventory.Update_Barang_Transaction_Inventory(Request, Request_kode)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else {
		result.Status = http.StatusNotFound
		result.Message = "Session Invalid"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}

func DeleteBarangTransaksiInventory(c echo.Context) error {
	var Request request.Update_Barang_Transaksi_Inventory_Kode_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")
	Request.Kode_barang_transaksi_inventory = c.Request().Header.Get("kode_barang_transaksi_inventory")

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {
		result, err = transaction_inventory.Delete_Barang_Transaksi_Inventory(Request)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

	} else {
		result.Status = http.StatusNotFound
		result.Message = "Session Invalid"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}

func UpdateStatusTransaksiInventory(c echo.Context) error {

	var Request request.Body_Update_Status_Transaksi_inventory
	var Request_session request.Token_Request
	var result response.Response
	var err error

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	_, condition := session_checking.Session_Checking(Request_session.Token)

	if condition {

		result, err = transaction_inventory.Update_Status_Transaksi_Inventory(Request)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else if !condition {
		result.Status = http.StatusGatewayTimeout
		result.Message = "Invalid Session"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}
