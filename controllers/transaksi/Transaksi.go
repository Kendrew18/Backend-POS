package transaksi

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/transaksi"
	"Bakend-POS/tools/session_checking"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputTransaksi(c echo.Context) error {
	var Request_body request.Body_Input_Transaksi_Request
	var Request request.Input_Transaksi_Request
	var Request_barang []request.Input_Barang_Transaksi_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request_body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	Request = Request_body.Input_transaksi
	Request_barang = Request_body.Input_barang_transaksi

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {

		result, err = transaksi.Input_Transaksi(Request, Request_barang)

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

func ReadTransaksi(c echo.Context) error {
	var Request request.Read_Transaksi_Request
	var Request_filter request.Read_Transaksi_Filter_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")
	Request_filter.Tanggal_awal = c.Request().Header.Get("tanggal_awal")
	Request_filter.Tanggal_akhir = c.Request().Header.Get("tanggal_akhir")
	Request_filter.Nama_customer = c.Request().Header.Get("nama_customer")

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {
		result, err = transaksi.Read_Transaksi(Request, Request_filter)

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
