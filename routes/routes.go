package routes

import (
	kasir "Bakend-POS/controllers/Kasir"
	"Bakend-POS/controllers/category"
	"Bakend-POS/controllers/foto"
	"Bakend-POS/controllers/home"
	"Bakend-POS/controllers/inventory"
	"Bakend-POS/controllers/jenis_pembayaran"
	"Bakend-POS/controllers/news"
	"Bakend-POS/controllers/transaction_invent"
	"Bakend-POS/controllers/transaksi"

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
		return c.String(http.StatusOK, "POS-V1")
	})

	US := e.Group("/US")
	SUP := e.Group("/SUP")
	INV := e.Group("/INV")
	TI := e.Group("/TI")
	KS := e.Group("/KS")
	FT := e.Group("/FT")
	TR := e.Group("/TR")
	HM := e.Group("/HM")
	JP := e.Group("/JP")
	CT := e.Group("/CT")
	NW := e.Group("/NW")

	//Jenis Pembayaran
	CT.GET("/category", category.ReadCategory)

	//User
	US.POST("/sign-up", user.SignUp)
	US.POST("/login", user.LoginUser)
	US.GET("/user-profile", user.UserProfile)
	US.PUT("/user-profile", user.UpdateUserProfile)
	US.PUT("/resend-otp", user.ResendOTP)
	US.PUT("/activate-acc", user.ActivateAccount)

	//Supplier
	SUP.POST("/supplier", supplier.InputSupplier)
	SUP.GET("/supplier", supplier.ReadSupplier)
	SUP.PUT("/supplier", supplier.UpdateSupplier)
	SUP.DELETE("/supplier", supplier.DeleteSupplier)

	//Stock
	INV.POST("/inventory", inventory.InputInventory)
	INV.GET("/inventory", inventory.ReadInventory)
	INV.PUT("/inventory", inventory.UpdateInventory)

	//Foto
	FT.GET("/foto", foto.ReadFoto)

	//Transaksi Inventory
	TI.POST("/transaction-inventory", transaction_invent.InputTransactionInventory)
	TI.GET("/transaction-inventory", transaction_invent.ReadTransactionInventory)
	TI.PUT("/update-header", transaction_invent.UpdateHeaderTransactionInventory)
	TI.PUT("/update-barang", transaction_invent.UpdateBarangTransactionInventory)
	TI.DELETE("/delete-barang", transaction_invent.DeleteBarangTransaksiInventory)
	TI.PUT("/update-status", transaction_invent.UpdateStatusTransaksiInventory)
	TI.GET("/dropdown-barang", transaction_invent.DropdownTransaksiInventory)

	//Kasir
	KS.GET("/kasir", kasir.ReadStockKasir)

	//Jenis Pembayaran
	JP.GET("/jenis-pembayaran", jenis_pembayaran.ReadJenisPembayaran)

	//Transksi
	TR.POST("/transaksi", transaksi.InputTransaksi)
	TR.GET("/transaksi", transaksi.ReadTransaksi)

	//Home
	HM.GET("/home", home.ReadHome)

	//News
	NW.POST("/news", news.InputNews)
	NW.GET("/news", news.ReadNewsAdmin)
	NW.GET("/news-user", news.ReadNewsUser)
	NW.DELETE("/news", news.DeleteNews)

	return e
}
