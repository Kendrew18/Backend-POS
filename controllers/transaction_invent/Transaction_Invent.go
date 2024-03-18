package transaction_invent

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/transaction_inventory"
	"Bakend-POS/tools/session_checking"
	"fmt"
	"net/http"
	"strings"

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

	var split []string

	//TOKEN,TANGGAL_AWAL,TANGGAL_AKHIR,NAMA_SUPPLIER
	Request_session.Token = c.Request().Header.Get("token")
	fmt.Println(Request_session.Token)
	split = strings.Split(Request_session.Token, ",")
	Request_session.Token = split[0]
	Request_filter.Tanggal_awal = split[1]
	Request_filter.Tanggal_akhir = split[2]
	Request_filter.Nama_supplier = split[3]

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

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request = Request_body.Update_header_transaksi_inventory
	Request_kode = Request_body.Update_header_transaksi_inventory_kode

	Request_kode.Kode_user = User.Kode_user

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

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request = Request_body.Update_barang_transaksi_inventory
	Request_kode = Request_body.Update_barang_transaksi_inventory_kode

	Request_kode.Kode_user = User.Kode_user

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
	kode_barang := c.Request().Header.Get("kode_barang")

	fmt.Println(kode_barang)

	Request.Kode_barang_transaksi_inventory = kode_barang
	fmt.Println(Request.Kode_barang_transaksi_inventory)

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
