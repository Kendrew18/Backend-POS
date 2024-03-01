package transaction_invent

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/transaction_inventory"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputTransactionInventory(c echo.Context) error {
	var Request request.Input_Transaksi_Body_Request

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := transaction_inventory.Input_Transaction_Inventory(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadTransactionInventory(c echo.Context) error {

	var Request request.Body_Read_Transaksi_Inventory_Request

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := transaction_inventory.Read_Transaction_Inventory(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateHeaderTransactionInventory(c echo.Context) error {

	var Request request.Body_Update_Header_Transaksi_Inventory_Request

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := transaction_inventory.Update_Header_Transaction_Inventory(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateBarangTransactionInventory(c echo.Context) error {

	var Request request.Body_Update_Barang_Transaksi_Inventory

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := transaction_inventory.Update_Barang_Transaction_Inventory(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
