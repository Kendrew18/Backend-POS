package foto

import "github.com/labstack/echo/v4"

func ReadFoto(c echo.Context) error {
	path := c.FormValue("path")
	return c.File(path)
}
