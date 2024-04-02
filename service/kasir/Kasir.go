package kasir

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"fmt"
	"net/http"
)

func Read_Stock_Kasir(Request request.Read_Kasir_Request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Kasir_Response

	con := db.CreateConGorm()

	err := con.Table("inventory").Select("kode_inventory", "nama_barang", "jumlah_barang", "satuan_barang", "harga_jual", "path_photo").Where("kode_user = ? && jumlah_barang > 0", Request.Kode_user).Order("co DESC").Scan(&arr_invent)
	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	for i := 0; i < len(arr_invent); i++ {
		var arr_detail_invent []response.Read_Detail_Inventory_Response
		err := con.Table("detail_inventory").Select("kode_barang_transaksi_inventory", "kode_inventory", "jumlah", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal").Joins("join transaksi_inventory ti on ti.kode_transaksi_inventory = detail_inventory.kode_transaksi_inventory").Where("kode_inventory = ? && jumlah > 0", arr_invent[i].Kode_inventory).Order("detail_inventory.co ASC").Scan(&arr_detail_invent)

		fmt.Println(arr_invent[i].Kode_inventory)

		fmt.Println(arr_detail_invent)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}
		arr_invent[i].Detail_inventory = arr_detail_invent
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}
