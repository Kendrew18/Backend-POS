package transaksi

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Input_Transaksi(Request request.Input_Transaksi_Request, Request_barang []request.Input_Barang_Transaksi_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	co := 0

	err := con.Table("transaksi").Select("co").Order("co DESC").Limit(1).Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	Request.Co = co + 1
	Request.Kode_transaksi = "TR-" + strconv.Itoa(Request.Co)

	date, _ := time.Parse("02-01-2006", Request.Tanggal)
	Request.Tanggal = date.Format("2006-01-02")
	Request.Kode_nota = date.Format("20060102") + "-"

	co_pembayaran := 0

	err = con.Table("transaksi").Select("COUNT(co)").Where("tanggal = ?", Request.Tanggal).Order("co DESC").Scan(&co_pembayaran)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	co_pembayaran = co_pembayaran + 1

	Request.Kode_nota = Request.Kode_nota + strconv.Itoa(co_pembayaran)
	Request.Jumlah_total = 0.0

	err = con.Table("transaksi").Select("co", "kode_transaksi", "kode_nota", "tanggal", "kode_jenis_pembayaran", "jumlah_total", "total_harga", "diskon", "tax", "kode_user", "nama_customer", "nomer_telp_customer", "alamat_customer").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	co = 0

	err = con.Table("barang_transaksi").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	total_harga := int64(0)
	total_jumlah := 0.0

	for i := 0; i < len(Request_barang); i++ {
		Request_barang[i].Co = co + 1 + i
		Request_barang[i].Kode_barang_transaksi = "BT-" + strconv.Itoa(Request_barang[i].Co)
		Request_barang[i].Kode_transaksi = Request.Kode_transaksi
		Request_barang[i].Sub_total = int64(math.Round(float64(Request_barang[i].Harga) * float64(Request_barang[i].Jumlah_barang)))
		total_harga = total_harga + Request_barang[i].Sub_total

		total_jumlah = total_jumlah + float64(Request_barang[i].Jumlah_barang)
	}

	total_harga = total_harga - Request.Diskon

	total_harga = total_harga + int64(math.Round(float64(total_harga)*Request.Tax/100))

	err = con.Table("barang_transaksi").Select("co", "kode_barang_transaksi", "kode_transaksi", "kode_barang_transaksi_inventory", "kode_inventory", "jumlah_barang", "harga", "nama_satuan", "sub_total").Create(&Request_barang)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("transaksi").Where("kode_transaksi = ?", Request.Kode_transaksi).Update("total_harga", total_harga).Update("jumlah_total", total_jumlah)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Exec("UPDATE `detail_inventory` JOIN barang_transaksi bsm ON bsm.kode_barang_transaksi_inventory = detail_inventory.kode_barang_transaksi_inventory SET `jumlah`= jumlah - jumlah_barang  WHERE bsm.kode_transaksi = ?", Request.Kode_transaksi)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Exec("UPDATE `inventory` JOIN barang_transaksi bsm ON bsm.kode_inventory = inventory.kode_inventory SET inventory.`jumlah_barang`=inventory.jumlah_barang - bsm.jumlah_barang WHERE bsm.kode_transaksi= ?", Request.Kode_transaksi)

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

func Read_Transaksi(Request request.Read_Transaksi_Request, Request_filter request.Read_Transaksi_Filter_Request) (response.Response, error) {

	var res response.Response
	var arr_data []response.Read_Transaksi_Response

	con := db.CreateConGorm()

	statement := "transaksi.kode_user = '" + Request.Kode_user + "'"

	fmt.Println(Request.Kode_user)

	if Request_filter.Nama_customer != "" {
		statement += " && transaksi.nama_customer = '" + Request_filter.Nama_customer + "'"
	}

	if Request_filter.Tanggal_awal != "" && Request_filter.Tanggal_akhir != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		date2, _ := time.Parse("02-01-2006", Request_filter.Tanggal_akhir)
		Request_filter.Tanggal_akhir = date2.Format("2006-01-02")

		statement += " AND (tanggal >= '" + Request_filter.Tanggal_awal + "' && tanggal <= '" + Request_filter.Tanggal_akhir + "' )"

	} else if Request_filter.Tanggal_awal != "" {

		date, _ := time.Parse("02-01-2006", Request_filter.Tanggal_awal)
		Request_filter.Tanggal_awal = date.Format("2006-01-02")

		statement += " && tanggal = '" + Request_filter.Tanggal_awal + "'"

	}

	fmt.Println(statement)

	err := con.Table("transaksi").Select("kode_transaksi", "DATE_FORMAT(tanggal, '%d-%m-%Y') AS tanggal", "kode_nota", "nama_customer", "nomer_telp_customer", "alamat_customer", "transaksi.kode_jenis_pembayaran", "nama_jenis_pembayaran", "jumlah_total", "total_harga", "tax", "diskon").Joins("join jenis_pembayaran jp on jp.kode_jenis_pembayaran = transaksi.kode_jenis_pembayaran").Where(statement).Order("transaksi.co DESC").Scan(&arr_data)

	fmt.Println(arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr_data
		return res, err.Error
	}

	for i := 0; i < len(arr_data); i++ {

		err = con.Table("barang_transaksi").Select("kode_barang_transaksi", "barang_transaksi.kode_inventory", "nama_barang", "barang_transaksi.jumlah_barang", "barang_transaksi.harga", "barang_transaksi.sub_total", "barang_transaksi.nama_satuan", "barang_transaksi.kode_barang_transaksi_inventory").Joins("join inventory s on s.kode_inventory = barang_transaksi.kode_inventory").Where("kode_transaksi = ?", arr_data[i].Kode_transaksi).Scan(&arr_data[i].Barang_transaksi)

		fmt.Println(arr_data)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			fmt.Println(err.Error)
			return res, err.Error
		}

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
