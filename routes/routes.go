package routes

import (
	"Bakend-POS/controllers/stock"
	"Bakend-POS/controllers/stock_masuk"
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
	STK := e.Group("/STK")
	STM := e.Group("/STM")
	// tr := e.Group("/tr")
	// pmb := e.Group("/pmb")
	// rtr := e.Group("/rtr")
	// flt := e.Group("/flt")

	//User
	US.GET("/login", user.LoginUser)
	US.GET("/user-profile", user.UserProfile)
	US.POST("/sign-up", user.SignUp)

	//Supplier
	SUP.POST("/supplier", supplier.InputSupplier)
	SUP.GET("/supplier", supplier.ReadSupplier)

	//Stock
	STK.POST("/stock", stock.InputStock)
	STK.GET("/stock", stock.ReadStock)
	STK.PUT("/stock", stock.UpdateStock)

	//Stock Masuk
	STM.POST("/stock-masuk", stock_masuk.InputStockMasuk)
	STM.GET("/stock-masuk", stock_masuk.ReadStockMasuk)
	//STM.PUT("/stock-masuk", stock_masuk.UpdateBarangStockMasuk)
	//STM.DELETE("/stock-masuk", stock_masuk.DeleteBarangStockMasuk)

	// //Input-Transaksi
	// tr.POST("/input-transaksi", controllers.InputTransaksi)

	// //read-transaksi
	// tr.GET("/transaksi", controllers.ReadTransaksi)

	// //read-detail-transaksi
	// tr.GET("/read-detail-transaksi", controllers.ReadDetailTransaksi)

	// //update-status-transaksi
	// tr.PUT("/update-status", controllers.UpdateStatus)

	// //tanggal-penjualan
	// tr.GET("/tgl-penjualan", controllers.DateTransaksi)

	// //penutupan-pembukuan
	// pmb.GET("/penutupan-pembukuan", controllers.PenutupanPembukuan)

	// //read-pembukuan
	// pmb.GET("/read-pembukuan", controllers.ReadPembukuan)

	// //input_retur
	// rtr.POST("/input-retur", controllers.InputRetur)

	// //read-retur
	// rtr.GET("/read-retur", controllers.ReadRetur)

	// //Read-Kode-Nama-Barang
	// rtr.GET("/read-kode-nama-barang", controllers.ReadKodeNamaBarang)

	// //Read-Max-Jumlah
	// rtr.GET("/read-max-jumlah", controllers.ReadMaxJumlah)

	// //filter_transaksi
	// flt.GET("/fil-transaksi", controllers.FilterTransaksi)

	// //filter_stock
	// flt.GET("/fil-stock", controllers.FilterStock)

	// //filter_stock_masuk
	// flt.GET("/fil-stock-masuk", controllers.FilterStockMasuk)

	// //filter_pembukuan
	// flt.GET("/fil-pembukuan", controllers.FilterReadPembukuan)

	// //read_filter_stock
	// flt.GET("/read-fil-stock", controllers.ReadFilterStock)

	return e
}
