package pembukuan

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools"
	"math"
	"net/http"
	"strconv"
	"time"
)

// func Read_Pembukuan(Request request.Read_Pembukuan_Request, Request_filter request.Read_Pembukuan_Filter_Request) (response.Response, error) {

// 	var res response.Response
// 	var arr_data []response.Read_Pembukuan_Response
// 	var data response.Read_Pembukuan_Response
// 	var rows *sql.Rows
// 	var err error

// 	con := db.CreateConGorm()

// 	statement := "stock_masuk.kode_user = '" + Request.Kode_user + "'"

// 	if Request_filter.Tanggal_awal != "" && Request_filter.Tanggal_akhir != "" {

// 		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
// 		Request_filter.Tanggal_awal = date.Format("2006-01-02")

// 		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_akhir)
// 		Request_filter.Tanggal_akhir = date2.Format("2006-01-02")

// 		statement += " AND (tanggal_masuk >= '" + Request_filter.Tanggal_awal + "' && tanggal_masuk <= '" + Request_filter.Tanggal_akhir + "' )"

// 	} else if Request_filter.Tanggal_awal != "" {

// 		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
// 		Request_filter.Tanggal_awal = date.Format("2006-01-02")

// 		statement += " && tanggal_masuk = '" + Request_filter.Tanggal_awal + "'"

// 	}

// 	rows, err = con.Table("pembukuan").Select("kode_pembukuan", "tanggal", "kode_nota", "pembukuan.kode_jenis_pembayaran", "nama_jenis_pembayaran", "diskon", "total_harga", "total_barang").Joins("JOIN jenis_pembayaran jp on jp.kode_jenis_pembayaran = pembukuan.kode_jenis_pembayaran").Where("kode_user=?", Request.Kode_user).Rows()

// 	if err != nil {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Status Not Found"
// 		res.Data = arr_data
// 		return res, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {

// 		err = rows.Scan(&data.Kode_pembukuan, &data.Tanggal, &data.Kode_nota, &data.Kode_jenis_pembayaran, &data.Nama_jenis_pembayaran, &data.Diskon,&data.Total_harga,&)

// 		if err != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Status Not Found"
// 			res.Data = data
// 			return res, err
// 		}

// 		var arr_barang []response.Read_Barang_Stock_Masuk_Response

// 		err = con.Table("barang_stock_masuk").Select("kode_barang_stock_masuk", "barang_stock_masuk.kode_stock", "nama_barang", "barang_stock_masuk.jumlah", "barang_stock_masuk.harga", "barang_stock_masuk.sub_total").Joins("join stock s on s.kode_stock = barang_stock_masuk.kode_stock").Where("kode_stock_masuk = ?", data.Kode_stock_masuk).Scan(&arr_barang).Error

// 		if err != nil {
// 			res.Status = http.StatusNotFound
// 			res.Message = "Status Not Found"
// 			res.Data = data
// 			return res, err
// 		}

// 		data.Barang_stock_masuk = arr_barang

// 		arr_data = append(arr_data, data)

// 	}

// 	if arr_data == nil {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Status Not Found"
// 		res.Data = arr_data

// 	} else {
// 		res.Status = http.StatusOK
// 		res.Message = "Suksess"
// 		res.Data = arr_data
// 	}
// 	return res, nil
// }

func Input_Pembukuan(Request request.Input_Pembukuan_Request, Request_barang request.Input_Barang_Pembukuan_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm()

	co := 0

	err := con.Table("pembukuan").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	Request.Co = co + 1
	Request.Kode_pembukuan = "BKD-" + strconv.Itoa(Request.Co)

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	err = con.Table("pembukuan").Select("co", "kode_pembukuan", "kode_nota", "tanggal", "kode_jenis_pembayaran", "diskon", "kode_user").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	kode_stock := tools.String_Separator_To_String(Request_barang.Kode_stock)
	jumlah_barang := tools.String_Separator_To_float64(Request_barang.Jumlah)
	harga_barang := tools.String_Separator_To_Int64(Request_barang.Harga)
	Satuan_barang := tools.String_Separator_To_String(Request_barang.Satuan_barang)

	co = 0

	err = con.Table("barang_pembukuan").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	var data request.Input_Barang_Pembukuan_V2_Request
	var arr_data []request.Input_Barang_Pembukuan_V2_Request

	for i := 0; i < len(kode_stock); i++ {
		data.Co = co + 1 + i
		data.Kode_barang_pembukuan = "BBK-" + strconv.Itoa(data.Co)
		data.Kode_pembukuan = Request.Kode_pembukuan
		data.Kode_stock = kode_stock[i]
		data.Jumlah = jumlah_barang[i]
		data.Harga = harga_barang[i]
		data.Satuan_barang = Satuan_barang[i]
		data.Sub_total = int64(math.Round(float64(harga_barang[i]) * jumlah_barang[i]))

		Request.Total_barang = Request.Total_barang + data.Jumlah
		Request.Total_harga = Request.Total_harga + data.Sub_total

		arr_data = append(arr_data, data)
	}

	err = con.Table("barang_pembukuan").Select("co", "kode_barang_pembukuan", "kode_pembukuan", "kode_stock", "jumlah", "harga", "satuan_barang", "sub_stotal").Create(&arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("pembukuan").Select("total_harga", "total_barang").Where("kode_pembukuan = ?", Request.Kode_pembukuan).Updates(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	//Update data
	err = con.Exec("UPDATE `stock` JOIN barang_pembukuan bsm ON bsm.kode_stock = stock.kode_stock SET `jumlah_barang`= jumlah_barang - jumlah WHERE bsm.kode_pembukuan = ?", Request.Kode_pembukuan)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

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
