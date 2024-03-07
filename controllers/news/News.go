package news

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/news"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputSupplier(c echo.Context) error {
	var Request request.Input_News_Request

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := news.Input_News(Request, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
