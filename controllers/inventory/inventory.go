package inventory

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/inventory"
	"Bakend-POS/tools/session_checking"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputInventory(c echo.Context) error {
	var Request request.Input_Inventory_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {

		result, err = inventory.Input_Inventory(Request)

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

func ReadInventory(c echo.Context) error {
	var Request request.Read_Inventory_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {
		result, err = inventory.Read_Inventory(Request)

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

func UpdateInventory(c echo.Context) error {
	var Request request.Update_Inventory_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {

		result, err = inventory.Update_Inventory(Request)

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

// func CheckNamaStock(c echo.Context) error {
// 	kode_stock := c.FormValue("kode_stock")
// 	nama_barang := c.FormValue("nama_barang")

// 	result, err := models.Check_Nama_Stock(kode_stock, nama_barang)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
