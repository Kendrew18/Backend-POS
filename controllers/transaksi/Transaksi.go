package transaksi

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/transaksi"
	"Bakend-POS/tools/session_checking"
	"fmt"
	"net/http"
	"strings"

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

	//TOKEN,TANGGAL_AWAL,TANGGAL_AKHIR,NAMA_CUSTOMER
	Request_session.Token = c.Request().Header.Get("token")
	fmt.Println(Request_session.Token)
	//fmt.Println(Request_session.Token)
	split := strings.Split(Request_session.Token, ",")
	Request_session.Token = split[0]
	Request_filter.Tanggal_awal = split[1]
	Request_filter.Tanggal_akhir = split[2]
	Request_filter.Nama_customer = split[3]

	fmt.Println(Request_session.Token)

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
