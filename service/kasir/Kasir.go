package kasir

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"net/http"
)

func Read_Stock_Kasir(Request request.Read_Kasir_Request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Kasir_Response

	con := db.CreateConGorm()

	err := con.Table("stock").Select("kode_stock", "nama_barang", "jumlah_barang", "satuan_barang", "harga_barang").Where("kode_user = ? && jumlah_barang > 0", Request.Kode_user).Order("co DESC").Scan(&arr_invent)
	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
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
