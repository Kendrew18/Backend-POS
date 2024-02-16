package routes

import (
	"Bakend-POS/controllers/inventory"
	"Bakend-POS/controllers/transaction_invent"

	"Bakend-POS/controllers/supplier"
	"Bakend-POS/controllers/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	US := e.Group("/US")
	SUP := e.Group("/SUP")
	INV := e.Group("/INV")
	TI := e.Group("/TI")
	// tr := e.Group("/tr")
	// pmb := e.Group("/pmb")
	// rtr := e.Group("/rtr")
	// flt := e.Group("/flt")

	//User
	US.POST("/sign-up", user.SignUp)
	US.GET("/login", user.LoginUser)
	US.GET("/user-profile", user.UserProfile)
	US.PUT("/user-profile", user.UpdateUserProfile)

	//Supplier
	SUP.POST("/supplier", supplier.InputSupplier)
	SUP.GET("/supplier", supplier.ReadSupplier)
	SUP.PUT("/supplier", supplier.UpdateSupplier)
	SUP.DELETE("/supplier", supplier.DeleteSupplier)

	//Stock
	INV.POST("/inventory", inventory.InputInventory)
	INV.GET("/inventory", inventory.ReadInventory)
	INV.PUT("/inventory", inventory.UpdateInventory)

	//Transaksi Inventory
	TI.POST("/transaction-inventory", transaction_invent.InputTransactionInventory)
	//TI.GET("/transaction-inventory", transaction_invent.ReadTransactionInventory)

	return e
}
