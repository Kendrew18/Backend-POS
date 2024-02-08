package supplier

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/supplier"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSupplier(c echo.Context) error {
	var Request request.Input_Supplier_Request

	Request.Email_supplier = c.FormValue("email_supplier")
	Request.Nama_supplier = c.FormValue("nama_supplier")
	Request.Nomor_telepon = c.FormValue("nomor_telepon")
	Request.Kode_user = c.FormValue("kode_user")

	result, err := supplier.Input_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadSupplier(c echo.Context) error {
	var Request request.Read_Supplier_Request

	Request.Kode_user = c.FormValue("kode_user")

	result, err := supplier.Read_Supplier(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

// func DeleteSupplier(c echo.Context) error {
// 	kode_supplier := c.FormValue("kode_supplier")

// 	result, err := models.Delete_Supplier(kode_supplier)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
