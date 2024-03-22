package category

import (
	"Bakend-POS/service/category"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadCategory(c echo.Context) error {
	result, err := category.Read_Category()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
