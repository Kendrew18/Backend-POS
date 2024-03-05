package kasir

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/kasir"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadStockKasir(c echo.Context) error {
	var Request request.Read_Kasir_Request

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := kasir.Read_Stock_Kasir(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
