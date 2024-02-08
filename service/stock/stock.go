package stock

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func Input_Stock(Request request.Input_Stock_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	nama_barang := ""

	err := con.Table("stock").Select("nama_barang").Where("nama_barang = ? AND kode_user = ?", Request.Nama_barang, Request.Kode_user).Scan(&nama_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if nama_barang == "" {

		co := 0

		err := con.Table("stock").Select("co").Order("co DESC").Scan(&co)

		Request.Co = co + 1
		Request.Kode_stock = "ST-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request.Jumlah_barang = math.Round(Request.Jumlah_barang*100) / 100

		err = con.Table("stock").Select("co", "kode_stock", "nama_barang", "jumlah_barang", "harga_barang", "satuan_barang", "kode_user").Create(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	return res, nil
}

func Read_Stock(Request request.Read_Stock_Request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Stock_Response

	con := db.CreateConGorm()

	err := con.Table("stock").Select("kode_stock", "nama_barang", "jumlah_barang", "satuan_barang", "harga_barang").Where("kode_user = ?", Request.Kode_user).Scan(&arr_invent)

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

func Update_Stock(Request request.Update_Stock_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm()

	kode_stock_masuk := ""

	err := con.Table("barang_stock_masuk").Select("kode_stock_masuk").Where("kode_stock = ?", Request.Kode_stock).Limit(1).Scan(&kode_stock_masuk)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	kode_barang_pembukuan := ""

	err = con.Table("barang_pembukuan").Select("kode_barang_pembukuan").Where("kode_stock = ?", Request.Kode_stock).Limit(1).Scan(&kode_barang_pembukuan)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if kode_stock_masuk == "" && kode_barang_pembukuan == "" {
		Request.Jumlah_barang = math.Round(Request.Jumlah_barang*100) / 100

		fmt.Println(Request)

		err := con.Table("stock").Where("kode_stock = ?", Request.Kode_stock).Select("nama_barang", "jumlah_barang", "harga_barang", "satuan_barang").Updates(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = map[string]int64{
				"rows": err.RowsAffected,
			}
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "data tidak dapat di update karena data telah terpakai"
		res.Data = Request
	}

	return res, nil
}

func Check_Nama_Stock(Request request.Check_Nama_Stock_Request) (response.Response, error) {
	var res response.Response
	var check response.Check_Nama_Stock_Response

	con := db.CreateConGorm()

	err := con.Table("stock").Select("nama_barang").Where("kode_user = ? AND Nama_barang = ? AND kode_stock != ", Request.Kode_user, Request.Nama_barang, Request.Kode_stock).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if check.Nama_barang == "" {
		check.Status = true
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = check
	} else {
		check.Status = false
		res.Status = http.StatusNotFound
		res.Message = "Nama Sudah Ada"
		res.Data = check
	}

	return res, nil
}
