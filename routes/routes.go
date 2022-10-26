package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/controllers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//Login
	e.GET("/login", controllers.Login)

	//User_Profile
	e.GET("/user-profile", controllers.User_Profile)

	//Input_Inventory
	e.POST("/input-inventory", controllers.InputInventory)

	//Read_Inventory
	e.GET("/inventory", controllers.ReadInventory)

	//Check_Nama_Inventory
	e.GET("/check-nama", controllers.Check_Nama_Inventory)

	return e
}
