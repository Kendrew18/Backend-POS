package inventory

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/inventory"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InputInventory(c echo.Context) error {
	var Request request.Input_Inventory_Request

	err := c.Bind(&Request)

	if err != nil {
		panic(err)
	}

	result, err := inventory.Input_Inventory(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadInventory(c echo.Context) error {
	var Request request.Read_Inventory_Request

	Request.Uuid_session = c.FormValue("uuid_session")

	result, err := inventory.Read_Inventory(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateInventory(c echo.Context) error {
	var Request request.Update_Inventory_Request

	Request.Kode_inventory = c.FormValue("kode_inventory")
	Request.Nama_barang = c.FormValue("nama_barang")
	Request.Harga_jual, _ = strconv.ParseInt(c.FormValue("harga_barang"), 10, 64)
	Request.Satuan_barang = c.FormValue("satuan_barang")
	Request.Uuid_session = c.FormValue("uuid_session")

	result, err := inventory.Update_Inventory(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

// func CheckNamaStock(c echo.Context) error {
// 	kode_stock := c.FormValue("kode_stock")
// 	nama_barang := c.FormValue("nama_barang")

// 	result, err := models.Check_Nama_Stock(kode_stock, nama_barang)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
