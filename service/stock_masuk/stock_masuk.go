package stock_masuk

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools"
	"database/sql"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Compare(data []request.Input_barang_Supplier_Request, kode_stock string) bool {

	for j := 0; j < len(data); j++ {

		if data[j].Kode_stock == kode_stock {
			return true
		}

	}

	return false
}

func Input_Stock_Masuk(Request request.Input_Stock_Masuk_Request, Request_barang request.Input_Barang_Stock_Masuk_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm()

	co := 0

	err := con.Table("stock_masuk").Select("co").Limit(1).Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_stock_masuk = "SM-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Tanggal_masuk)
	Request.Tanggal_masuk = date.Format("2006-01-02")

	err = con.Table("stock_masuk").Select("co", "kode_stock_masuk", "nama_penanggung_jawab", "tanggal_masuk", "kode_supplier", "kode_user").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	kode_stock := tools.String_Separator_To_String(Request_barang.Kode_stock)
	jumlah_barang := tools.String_Separator_To_float64(Request_barang.Jumlah)
	harga_barang := tools.String_Separator_To_Int64(Request_barang.Harga)

	co = 0

	err = con.Table("barang_stock_masuk").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	var data request.Input_Barang_Stock_Masuk_V2_Request
	var arr_data []request.Input_Barang_Stock_Masuk_V2_Request

	for i := 0; i < len(kode_stock); i++ {
		data.Co = co + 1 + i
		data.Kode_barang_stock_masuk = "BSM-" + strconv.Itoa(data.Co)
		data.Kode_stock_masuk = Request.Kode_stock_masuk
		data.Kode_stock = kode_stock[i]

		data.Jumlah = jumlah_barang[i]
		data.Harga = harga_barang[i]
		data.Sub_total = int64(math.Round(float64(harga_barang[i]) * jumlah_barang[i]))

		arr_data = append(arr_data, data)
	}

	err = con.Table("barang_stock_masuk").Select("co", "kode_barang_stock_masuk", "kode_stock_masuk", "kode_stock", "jumlah", "harga", "sub_total").Create(&arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	// add barang supplier yang belum masuk

	var data_sup request.Input_barang_Supplier_Request
	var arr_data_sup []request.Input_barang_Supplier_Request

	co = 0

	err = con.Table("barang_supplier").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	var data_stock []request.Input_barang_Supplier_Request

	err = con.Table("barang_supplier").Select("kode_stock").Joins("JOIN supplier s on s.kode_supplier = barang_supplier.kode_supplier").Where("barang_supplier.kode_supplier = ?", Request.Kode_supplier).Scan(&data_stock)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	x := 0

	for i := 0; i < len(kode_stock); i++ {
		if Compare(data_stock, kode_stock[i]) {

		} else {
			data_sup.Co = co + 1 + x
			data_sup.Kode_barang_supplier = "BS-" + strconv.Itoa(data.Co)
			data_sup.Kode_stock = kode_stock[i]
			data_sup.Kode_supplier = Request.Kode_supplier

			arr_data_sup = append(arr_data_sup, data_sup)
			x++
		}
	}

	if len(arr_data_sup) > 0 {
		err = con.Table("barang_supplier").Select("co", "kode_barang_supplier", "kode_supplier", "kode_stock").Create(&arr_data_sup)
	}

	//Update data
	err = con.Exec("UPDATE `stock` JOIN barang_stock_masuk bsm ON bsm.kode_stock = stock.kode_stock SET `jumlah_barang`=jumlah_barang + jumlah WHERE bsm.kode_stock_masuk = ?", Request.Kode_stock_masuk)

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

func Read_Stock_Masuk(Request request.Read_Stock_Masuk_Request, Request_filter request.Read_Stock_Masuk_Filter_Request) (response.Response, error) {
	var res response.Response
	var arr_data []response.Read_Stock_Masuk_Response
	var data response.Read_Stock_Masuk_Response
	var rows *sql.Rows
	var err error

	con := db.CreateConGorm()

	statement := "SELECT stock_masuk.kode_stock_masuk, DATE_FORMAT(tanggal_masuk, '%d-%m-%Y') AS tanggal_masuk, stock_masuk.kode_supplier, nama_supplier, sum(jumlah), sum(sub_total) FROM stock_masuk JOIN barang_stock_masuk bsm on bsm.kode_stock_masuk = stock_masuk.kode_stock_masuk JOIN supplier sp ON sp.kode_supplier = stock_masuk.kode_supplier WHERE stock_masuk.kode_user = '" + Request.Kode_user + "'"

	if Request_filter.Kode_supplier != "" {
		statement += " && stock_masuk.kode_supplier = '" + Request_filter.Kode_supplier + "'"
	}

	if Request_filter.Tanggal_awal != "" && Request_filter.Tanggal_akhir != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_akhir)
		Request_filter.Tanggal_akhir = date2.Format("2006-01-02")

		statement += " AND (tanggal_masuk >= '" + Request_filter.Tanggal_awal + "' && tanggal_masuk <= '" + Request_filter.Tanggal_akhir + "' )"

	} else if Request_filter.Tanggal_awal != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		statement += " && tanggal_masuk = '" + Request_filter.Tanggal_awal + "'"

	}

	statement += " GROUP BY stock_masuk.kode_stock_masuk ORDER BY stock_masuk.co DESC"

	rows, err = con.Raw(statement).Rows()

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Kode_stock_masuk, &data.Tanggal_masuk, &data.Kode_supplier, &data.Nama_supplier, &data.Total_barang, &data.Total_harga)

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		var arr_barang []response.Read_Barang_Stock_Masuk_Response

		err = con.Table("barang_stock_masuk").Select("kode_barang_stock_masuk", "barang_stock_masuk.kode_stock", "nama_barang", "barang_stock_masuk.jumlah", "barang_stock_masuk.harga", "barang_stock_masuk.sub_total").Joins("join stock s on s.kode_stock = barang_stock_masuk.kode_stock").Where("kode_stock_masuk = ?", data.Kode_stock_masuk).Scan(&arr_barang).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data
			return res, err
		}

		data.Barang_stock_masuk = arr_barang

		arr_data = append(arr_data, data)

	}

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data

	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = arr_data
	}
	return res, nil
}

// func Update_Barang_Stock_Masuk(Request request.Update_Stock_Masuk_Request, Request_kode request.Update_Kode_Barang_Stock_Masuk_Request) (response.Response, error) {
// 	var res response.Response

// 	check := -1
// 	con_check := db.CreateConGorm().Table("stock_masuk")

// 	err := con_check.Select("status").Joins("JOIN barang_stock_masuk bsm ON bsm.kode_stock_masuk = stock_masuk.kode_stock_masuk ").Where("kode_barang_stock_masuk = ?", Request_kode.Kode_barang_stock_masuk).Scan(&check)

// 	fmt.Println(check)

// 	if err.Error != nil {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Update Error"
// 		res.Data = Request
// 		return res, err.Error
// 	}

// 	if check == 0 || check == 2 {

// 		con := db.CreateConGorm().Table("barang_stock_masuk")

// 		Request.Sub_total = int64(math.Round(float64(Request.Harga) * Request.Jumlah))

// 		err = con.Where("kode_barang_stock_masuk = ?", Request_kode.Kode_barang_stock_masuk).Select("kode_stock", "harga", "jumlah", "sub_total").Updates(&Request)

// 		if err.Error != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Status Not Found"
// 			res.Data = Request
// 			return res, err.Error
// 		} else {
// 			res.Status = http.StatusOK
// 			res.Message = "Suksess"
// 			res.Data = map[string]int64{
// 				"rows": err.RowsAffected,
// 			}
// 		}
// 	} else {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Barang Tidak dapat di update"
// 		res.Data = Request
// 		return res, err.Error
// 	}
// 	return res, nil
// }

// func Delete_Barang_Stock_Masuk(Request request.Update_Kode_Barang_Stock_Masuk_Request) (response.Response, error) {
// 	var res response.Response

// 	check := -1
// 	con_check := db.CreateConGorm().Table("stock_masuk")

// 	err := con_check.Select("status").Joins("JOIN barang_stock_masuk bsm ON bsm.kode_stock_masuk = stock_masuk.kode_stock_masuk ").Where("kode_barang_stock_masuk = ?", Request.Kode_barang_stock_masuk).Scan(&check)

// 	fmt.Println(check)

// 	if err.Error != nil {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Update Error"
// 		res.Data = Request
// 		return res, err.Error
// 	}

// 	if check == 0 || check == 2 {

// 		con := db.CreateConGorm().Table("stock_masuk")

// 		data := ""

// 		err = con.Select("stock_masuk.kode_stock_masuk").Joins("JOIN barang_stock_masuk bsm ON bsm.kode_stock_masuk = stock_masuk.kode_stock_masuk").Where("kode_barang_stock_masuk = ?", Request.Kode_barang_stock_masuk).Scan(&data)

// 		if err.Error != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Update Error"
// 			res.Data = Request
// 			return res, err.Error
// 		}

// 		con_barang := db.CreateConGorm().Table("barang_stock_masuk")

// 		err = con_barang.Where("kode_barang_stock_masuk = ?", Request.Kode_barang_stock_masuk).Delete("")

// 		if err.Error != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Delete Error"
// 			res.Data = Request
// 			return res, err.Error
// 		}

// 		kode_barang := ""

// 		con_check := db.CreateConGorm().Table("barang_stock_masuk")

// 		err = con_check.Select("kode_barang_stock_masuk").Where("kode_stock_masuk=?", data).Limit(1).Scan(&kode_barang)

// 		if err.Error != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Delete Error"
// 			res.Data = Request
// 			return res, err.Error
// 		}

// 		if kode_barang == "" {

// 			con_del_req := db.CreateConGorm().Table("stock_masuk")

// 			err = con_del_req.Where("kode_stock_masuk = ?", data).Delete("")

// 			if err.Error != nil {
// 				res.Status = http.StatusNotFound
// 				res.Message = "Status Not Found"
// 				res.Data = Request
// 				return res, err.Error
// 			}
// 		}

// 		if err.Error != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Status Not Found"
// 			res.Data = Request
// 			return res, err.Error
// 		} else {
// 			res.Status = http.StatusOK
// 			res.Message = "Suksess"
// 			res.Data = map[string]int64{
// 				"rows": err.RowsAffected,
// 			}
// 		}
// 	} else {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Barang Tidak dapat di update"
// 		res.Data = Request
// 		return res, err.Error
// 	}
// 	return res, nil
// }

// func Update_Status_Stock_Masuk(Request request.Update_Status_Stock_Masuk_Request, Request_kode request.Update_Kode_Stock_Masuk_Request) (response.Response, error) {
// 	var res response.Response
// 	var err2 error
// 	con := db.CreateConGorm().Table("pre_order")
// 	status := -1

// 	err := con.Select("status").Where("kode_pre_order = ?", Request_kode).Scan(&status)

// 	if err.Error != nil {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Status Not Found"
// 		res.Data = Request
// 		return res, err.Error
// 	}

// 	if status != 1 {
// 		if Request.Status == 2 || Request.Status == 0 {

// 			con := db.CreateConGorm().Table("stock_masuk")

// 			err := con.Where("kode_stock_masuk = ?", Request_kode.Kode_stock_masuk).Select("status").Updates(&Request)

// 			if err.Error != nil {
// 				res.Status = http.StatusNotFound
// 				res.Message = "Status Not Found"
// 				res.Data = Request
// 				return res, err.Error
// 			} else {
// 				res.Status = http.StatusOK
// 				res.Message = "Suksess"
// 				res.Data = map[string]int64{
// 					"rows": err.RowsAffected,
// 				}
// 			}
// 		} else if Request.Status == 1 {
// 			con := db.CreateConGorm().Table("stock_masuk")

// 			err := con.Where("kode_stock_masuk = ?", Request_kode.Kode_stock_masuk).Select("status").Updates(&Request)

// 			// add barang supplier yang belum masuk

// 			var data_sup request.Input_barang_Supplier_Request
// 			var arr_data_sup []request.Input_barang_Supplier_Request

// 			co := 0

// 			err = con.Table("barang_supplier").Select("co").Order("co DESC").Scan(&co)

// 			if err.Error != nil {
// 				res.Status = http.StatusNotFound
// 				res.Message = "Status Not Found"
// 				res.Data = Request
// 				return res, err.Error
// 			}

// 			var data_stock []request.Input_barang_Supplier_Request

// 			err = con.Table("barang_supplier").Select("kode_stock").Joins("JOIN supplier s on s.kode_supplier = barang_supplier.kode_supplier").Where("barang_supplier.kode_supplier = ?", Request.Kode_supplier).Scan(&data_stock)

// 			if err.Error != nil {
// 				res.Status = http.StatusNotFound
// 				res.Message = "Status Not Found"
// 				res.Data = Request
// 				return res, err.Error
// 			}

// 			x := 0

// 			for i := 0; i < len(kode_stock); i++ {
// 				if Compare(data_stock, kode_stock[i]) {

// 				} else {
// 					data_sup.Co = co + 1 + x
// 					data_sup.Kode_barang_supplier = "BS-" + strconv.Itoa(data.Co)
// 					data_sup.Kode_stock = kode_stock[i]
// 					data_sup.Kode_supplier = Request.Kode_supplier

// 					arr_data_sup = append(arr_data_sup, data_sup)
// 					x++
// 				}
// 			}

// 			err = con.Table("barang_supplier").Select("co", "kode_barang_supplier", "kode_supplier", "kode_stock").Create(&arr_data_sup)

// 			//read data stock
// 			var stock []response.Read_Stock_Response

// 			err = con.Table("stock").Select("kode_stock", "(jumlah_stock + jumlah) as jumlah_stock ").Joins("JOIN barang_stock_masuk bsm ON bsm.kode_stock = stock.kode_stock").Where("kode_user = ? && kode_stock_masuk = ? && "+statement_in, Request.Kode_stock_masuk, Request.Kode_user).Scan(&stock)

// 			if err.Error != nil {
// 				res.Status = http.StatusNotFound
// 				res.Message = "Status Not Found"
// 				res.Data = Request
// 				return res, err.Error
// 			}

// 			err = con.Table("stock").Select("jumlah_stock").Where(statement_in).Updates(&arr_data)

// 			if err.Error != nil {
// 				res.Status = http.StatusNotFound
// 				res.Message = "Status Not Found"
// 				res.Data = Request
// 				return res, err.Error
// 			}

// 			if err2 != nil {
// 				res.Status = http.StatusNotFound
// 				res.Message = "Status Not Found"
// 				res.Data = Request
// 				return res, err.Error
// 			} else {
// 				res.Status = http.StatusOK
// 				res.Message = "Suksess"
// 				res.Data = map[string]int64{
// 					"rows": err.RowsAffected,
// 				}
// 			}

// 		}
// 	} else {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Tidah dapat di edit diakrenakan sudah sukses"
// 		res.Data = Request
// 	}
// 	return res, nil
// }
