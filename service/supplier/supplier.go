package supplier

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
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

	con := db.CreateConGorm().Table("supplier")

	err := con.Select("kode_supplier", "nama_supplier", "email_supplier", "nomor_telepon").Where("kode_user = ?", Request.Kode_user).Order("co ASC").Scan(&data).Error

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

	return res, nil
}

//func Delete_Supplier(kode_supplier string) (Response, error) {
// 	var res Response
// 	var arrobj []str.Read_Kode_Stock_Masuk
// 	var obj str.Read_Kode_Stock_Masuk

// 	con := db.CreateCon()

// 	sqlStatement := "SELECT id_stock_masuk,kode_supplier FROM stock_masuk WHERE kode_supplier=? "

// 	rows, err := con.Query(sqlStatement, kode_supplier)

// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&obj.Id_stock_masuk, &obj.Kode_supplier)
// 		if err != nil {
// 			return res, err
// 		}
// 		arrobj = append(arrobj, obj)
// 	}

// 	if arrobj == nil {

// 		sqlstatement := "DELETE FROM supplier WHERE kode_supplier=?"

// 		stmt, err := con.Prepare(sqlstatement)

// 		if err != nil {
// 			return res, err
// 		}

// 		result, err := stmt.Exec(kode_supplier)

// 		if err != nil {
// 			return res, err
// 		}

// 		rowsAffected, err := result.RowsAffected()

// 		if err != nil {
// 			return res, err
// 		}

// 		res.Status = http.StatusOK
// 		res.Message = "Suksess"
// 		res.Data = map[string]int64{
// 			"rows": rowsAffected,
// 		}

// 	} else {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Tidak bisa di hapus"
// 		res.Data = arrobj
// 	}

// 	return res, nil
// }
