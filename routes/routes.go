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

	Sup := e.Group("/sup")
	us := e.Group("/us")
	invent := e.Group("/invent")
	stk_m := e.Group("/stk-m")

	//Login
	us.GET("/login", controllers.Login)

	//User_Profile
	us.GET("/user-profile", controllers.User_Profile)

	//Input_Inventory
	invent.POST("/input-stock", controllers.InputStock)

	//Read_Inventory
	invent.GET("/stock", controllers.ReadStock)

	//Update_Inventory
	invent.PUT("/update-stock", controllers.UpdateStock)

	//Check_Nama_Inventory
	invent.GET("/check-nama", controllers.CheckNamaStock)

	//Input_Supplier
	Sup.POST("/input-supplier", controllers.InputSupplier)

	//Read_Supplier
	Sup.GET("/supplier", controllers.ReadSupplier)

	//Input_Stock_Masuk
	stk_m.POST("/input-stock-masuk", controllers.InputStockMasuk)

	//Read_Stock_Masuk
	stk_m.GET("/stock-masuk", controllers.ReadStockMasuk)

	//Read_Detail_Stock_Masuk
	stk_m.GET("/detail-stock-masuk", controllers.Read_Detail_Stock_Masuk)

	return e
}
