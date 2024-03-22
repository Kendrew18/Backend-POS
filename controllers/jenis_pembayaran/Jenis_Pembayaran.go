package jenis_pembayaran

import (
	"Bakend-POS/service/jenis_pembayaran"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadJenisPembayaran(c echo.Context) error {
	result, err := jenis_pembayaran.Read_Jenis_Pembayaran()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
