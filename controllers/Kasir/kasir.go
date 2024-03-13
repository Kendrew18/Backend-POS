package kasir

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/kasir"
	"Bakend-POS/tools/session_checking"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadStockKasir(c echo.Context) error {
	var Request request.Read_Kasir_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {

		result, err = kasir.Read_Stock_Kasir(Request)

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
