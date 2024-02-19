package supplier

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/session_checking"
	"net/http"
	"strconv"
)

func Input_Supplier(Request request.Input_Supplier_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm().Table("supplier")

	co := 0

	err := con.Select("co").Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_supplier = "SP-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Select("co", "kode_supplier", "nama_supplier", "nomor_telepon", "email_supplier", "kode_user").Create(&Request)

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

	return res, nil
}

func Read_Supplier(Request request.Read_Supplier_Request) (response.Response, error) {
	var res response.Response
	var data []response.Read_Supplier_Response

	User, condition := session_checking.Session_Checking(Request.Uuid_session)

	if condition {

		con := db.CreateConGorm().Table("supplier")

		err := con.Select("kode_supplier", "nama_supplier", "email_supplier", "nomor_telepon").Where("kode_user = ?", User.Kode_user).Order("co ASC").Scan(&data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err
		}

		if data == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = data
		} else {
			res.Status = http.StatusOK
			res.Message = "Sukses"
			res.Data = data
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Session Invalid"
		res.Data = data
	}

	return res, nil
}

func Update_Supplier(Request request.Update_Supplier_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm()

	err := con.Table("supplier").Where("kode_supplier = ?", Request.Kode_supplier).Select("nomor_telepon", "email_supplier").Updates(&Request)

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

	return res, nil
}

func Delete_Supplier(Request request.Delete_Supplier_Request) (response.Response, error) {
	var res response.Response

	check := ""
	con := db.CreateConGorm()

	err := con.Table("barang_supplier").Select("kode_barang_supplier").Where("kode_supplier = ?", Request.Kode_supplier).Limit(1).Scan(&check)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Update Error"
		res.Data = Request
		return res, err.Error
	}

	if check == "" {

		err = con.Table("supplier").Where("kode_supplier = ?", Request.Kode_supplier).Delete("")

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
		res.Message = "Barang Tidak dapat di update"
		res.Data = Request
		return res, err.Error
	}
	return res, nil
}
