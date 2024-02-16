package transaction_inventory

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Compare(data []request.Input_barang_Supplier_Request, kode_inventory string) bool {

	for j := 0; j < len(data); j++ {

		if data[j].Kode_inventory == kode_inventory {
			return true
		}

	}

	return false
}

func Input_Transaction_Inventory(Request request.Input_Transaksi_Inventory_Request, Request_barang request.Input_Barang_Transaksi_Inventory_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm()

	co := 0

	err := con.Table("transaksi_inventory").Select("co").Limit(1).Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_transaksi_inventory = "TI-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")

	err = con.Table("transaksi_inventory").Select("co", "kode_transaksi_inventory", "tanggal", "kode_transaksi", "kode_jenis_pembayaran", "harga_ongkos_kirim", "ppn", "kode_supplier", "kode_user", "jenis_transaksi").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	kode_inventory := tools.String_Separator_To_String(Request_barang.Kode_inventory)
	jumlah_barang := tools.String_Separator_To_float64(Request_barang.Jumlah)
	harga_barang := tools.String_Separator_To_Int64(Request_barang.Harga)

	co = 0

	err = con.Table("barang_transaksi_inventory").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	var data request.Input_Barang_Transaksi_Inventory_V2_Request
	var arr_data []request.Input_Barang_Transaksi_Inventory_V2_Request

	total_harga := int64(0)
	total_jumlah := 0.0

	for i := 0; i < len(kode_inventory); i++ {
		data.Co = co + 1 + i
		data.Kode_barang_transaksi_inventory = "BTI-" + strconv.Itoa(data.Co)
		data.Kode_transaksi_inventory = Request.Kode_transaksi_inventory
		data.Kode_inventory = kode_inventory[i]
		data.Jumlah = jumlah_barang[i]
		data.Harga = harga_barang[i]
		data.Sub_total = int64(math.Round(float64(harga_barang[i]) * jumlah_barang[i]))

		total_harga = total_harga + data.Sub_total
		total_jumlah = total_jumlah + data.Jumlah

		arr_data = append(arr_data, data)
	}

	fmt.Println(arr_data)

	total_harga = total_harga + Request.Harga_ongkos_kirim + int64(math.Round(float64(total_harga)*Request.Ppn/100))

	err = con.Table("barang_transaksi_inventory").Select("co", "kode_barang_transaksi_inventory", "kode_transaksi_inventory", "kode_inventory", "jumlah", "harga", "sub_total").Create(&arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("transaksi_inventory").Where("kode_transaksi_inventory = ?", Request.Kode_transaksi_inventory).Update("total_harga", total_harga).Update("total_barang", total_jumlah)

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

	err = con.Table("barang_supplier").Select("kode_inventory").Joins("JOIN supplier s on s.kode_supplier = barang_supplier.kode_supplier").Where("barang_supplier.kode_supplier = ?", Request.Kode_supplier).Scan(&data_stock)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	x := 0

	for i := 0; i < len(kode_inventory); i++ {
		if Compare(data_stock, kode_inventory[i]) {

		} else {
			data_sup.Co = co + 1 + x
			data_sup.Kode_barang_supplier = "BS-" + strconv.Itoa(data.Co)
			data_sup.Kode_inventory = kode_inventory[i]
			data_sup.Kode_supplier = Request.Kode_supplier

			arr_data_sup = append(arr_data_sup, data_sup)
			x++
		}
	}

	if len(arr_data_sup) > 0 {
		err = con.Table("barang_supplier").Select("co", "kode_barang_supplier", "kode_supplier", "kode_inventory").Create(&arr_data_sup)
		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}
	}

	//Update data
	err = con.Exec("UPDATE `inventory` JOIN barang_transaksi_inventory bsm ON bsm.kode_inventory = inventory.kode_inventory SET `jumlah_barang`=jumlah_barang + jumlah WHERE bsm.kode_transaksi_inventory = ?", Request.Kode_transaksi_inventory)

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

// func Read_Transaction_Inventory(Request request.Read_Transaksi_Inventory_Request, Request_filter request.Read_Stock_Masuk_Filter_Request) (response.Response, error) {
// 	var res response.Response
// 	var arr_data []response.Read_Stock_Masuk_Response
// 	var data response.Read_Stock_Masuk_Response
// 	var rows *sql.Rows
// 	var err error

// 	con := db.CreateConGorm()

// 	statement := "SELECT stock_masuk.kode_stock_masuk, DATE_FORMAT(tanggal_masuk, '%d-%m-%Y') AS tanggal_masuk, stock_masuk.kode_supplier, nama_supplier, sum(jumlah), sum(sub_total) FROM stock_masuk JOIN barang_stock_masuk bsm on bsm.kode_stock_masuk = stock_masuk.kode_stock_masuk JOIN supplier sp ON sp.kode_supplier = stock_masuk.kode_supplier WHERE stock_masuk.kode_user = '" + Request.Kode_user + "'"

// 	if Request_filter.Kode_supplier != "" {
// 		statement += " && stock_masuk.kode_supplier = '" + Request_filter.Kode_supplier + "'"
// 	}

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

// 	statement += " GROUP BY stock_masuk.kode_stock_masuk ORDER BY stock_masuk.co DESC"

// 	rows, err = con.Raw(statement).Rows()

// 	if err != nil {
// 		res.Status = http.StatusNotFound
// 		res.Message = "Status Not Found"
// 		res.Data = arr_data
// 		return res, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {

// 		err = rows.Scan(&data.Kode_stock_masuk, &data.Tanggal_masuk, &data.Kode_supplier, &data.Nama_supplier, &data.Total_barang, &data.Total_harga)

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
