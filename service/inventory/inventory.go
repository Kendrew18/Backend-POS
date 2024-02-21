package inventory

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/session_checking"
	"fmt"
	"net/http"
	"strconv"
)

func Input_Inventory(Request request.Input_Inventory_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	nama_barang := ""

	err := con.Table("inventory").Select("nama_barang").Where("nama_barang = ? AND kode_user = ?", Request.Nama_barang, Request.Kode_user).Scan(&nama_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	User, condition := session_checking.Session_Checking(Request.Uuid_session)

	if nama_barang == "" && condition {

		Request.Kode_user = User.Kode_user

		co := 0

		err := con.Table("inventory").Select("co").Order("co DESC").Scan(&co)

		Request.Co = co + 1
		Request.Kode_inventory = "IN-" + strconv.Itoa(Request.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("inventory").Select("co", "kode_inventory", "nama_barang", "harga_jual", "satuan_barang", "kode_user").Create(&Request)

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

	} else if !condition {
		res.Status = http.StatusNotFound
		res.Message = "Session Invalid"
		res.Data = Request
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Nama Barang Telah Digunakan"
		res.Data = Request
		return res, err.Error
	}

	return res, nil
}

func Read_Inventory(Request request.Read_Inventory_Request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Inventory_Response

	User, condition := session_checking.Session_Checking(Request.Uuid_session)

	if condition {
		con := db.CreateConGorm()

		err := con.Table("inventory").Select("kode_inventory", "nama_barang", "jumlah_barang", "satuan_barang", "harga_jual").Where("kode_user = ?", User.Kode_user).Scan(&arr_invent)

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Nama Barang Telah Digunakan"
		res.Data = Request
	}

	return res, nil
}

func Update_Inventory(Request request.Update_Inventory_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	kode_inventory := ""

	fmt.Println(Request)

	err := con.Table("barang_supplier").Select("kode_inventory").Where("kode_inventory = ?", Request.Kode_inventory).Limit(1).Scan(&kode_inventory)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	// kode_barang_pembukuan := ""

	// err = con.Table("barang_pembukuan").Select("kode_barang_pembukuan").Where("kode_stock = ?", Request.Kode_stock).Limit(1).Scan(&kode_barang_pembukuan)

	// if err.Error != nil {
	// 	res.Status = http.StatusNotFound
	// 	res.Message = "Status Not Found"
	// 	res.Data = Request
	// 	return res, err.Error
	// }

	_, condition := session_checking.Session_Checking(Request.Uuid_session)

	if kode_inventory == "" && condition {

		err := con.Table("inventory").Where("kode_inventory = ?", Request.Kode_inventory).Select("nama_barang", "harga_jual", "satuan_barang").Updates(&Request)

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
	} else if !condition {
		res.Status = http.StatusNotFound
		res.Message = "Session Invalid"
		res.Data = Request
	} else {
		res.Status = http.StatusNotFound
		res.Message = "data tidak dapat di update karena data telah terpakai"
		res.Data = Request

		return res, err.Error
	}

	return res, nil
}

func Check_Nama_Inventory(Request request.Check_Nama_Inventory_Request) (response.Response, error) {
	var res response.Response
	var check response.Check_Nama_Inventory_Response

	con := db.CreateConGorm()

	err := con.Table("inventory").Select("nama_barang").Where("kode_user = ? AND nama_barang = ? AND kode_inventory != ", Request.Kode_user, Request.Nama_barang, Request.Kode_inventory).Scan(&check)

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
