package supplier

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/supplier"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BodyStuct struct {
	Request   request.Input_Supplier_Request `json:"json1"`
	Request_2 []request.Read_News_Request    `json:"json2"`
}

func InputSupplier(c echo.Context) error {
	var Request BodyStuct

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	fmt.Println(Request)

	result, err := supplier.Input_Supplier(Request.Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadSupplier(c echo.Context) error {
	var Request request.Read_Supplier_Request

	Request.Uuid_session = c.FormValue("uuid_session")

	result, err := supplier.Read_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateSupplier(c echo.Context) error {
	var Request request.Update_Supplier_Request

	Request.Kode_supplier = c.FormValue("kode_supplier")
	Request.Email_supplier = c.FormValue("email_supplier")
	Request.Nomor_telepon = c.FormValue("nomor_telepon")
	Request.Uuid_session = c.FormValue("uuid_session")

	result, err := supplier.Update_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteSupplier(c echo.Context) error {
	var Request request.Delete_Supplier_Request
	Request.Kode_supplier = c.FormValue("kode_supplier")

	result, err := supplier.Delete_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
