package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"project-1/controllers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	Sup := e.Group("/sup")
	us := e.Group("/us")
	invent := e.Group("/invent")
	stk_m := e.Group("/stk-m")
	tr := e.Group("/tr")
	pmb := e.Group("/pmb")

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

	//Delete_Supplier
	Sup.DELETE("/delete-supplier", controllers.DeleteSupplier)

	//Input_Stock_Masuk
	stk_m.POST("/input-stock-masuk", controllers.InputStockMasuk)

	//Read_Stock_Masuk
	stk_m.GET("/stock-masuk", controllers.ReadStockMasuk)

	//Read_Detail_Stock_Masuk
	stk_m.GET("/detail-stock-masuk", controllers.Read_Detail_Stock_Masuk)

	//Input-Transaksi
	tr.POST("/input-transaksi", controllers.InputTransaksi)

	//read-transaksi
	tr.GET("/transaksi", controllers.ReadTransaksi)

	//read-detail-transaksi
	tr.GET("/read-detail-transaksi", controllers.ReadDetailTransaksi)

	//update-status-transaksi
	tr.PUT("/update-status", controllers.UpdateStatus)

	//penutupan-pembukuan
	pmb.GET("/penutupan-pembukuan", controllers.PenutupanPembukuan)

	pmb.GET("/read-pembukuan", controllers.PenutupanPembukuan)

	return e
}
