package home

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Read_Home(Request request.Home_Request) (response.Response, error) {
	var res response.Response
	var arr_invent response.Home_Response

	con := db.CreateConGorm()

	tm := time.Now()
	//tanggal := tm.Format("2006-01-02")
	year := tm.Format("2006")

	fmt.Println(year)

	err := con.Table("transaksi_inventory").Select("IFNULL(SUM(total_harga),0) as total_pengeluaran").Where("kode_user = ? && DATE_FORMAT(tanggal, '%Y') = ? && status = 1", Request.Kode_user, year).Scan(&arr_invent.Total_pengeluaran)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("transaksi").Select("IFNULL(SUM(total_harga),0) as total_pemasukan").Where("kode_user = ? && DATE_FORMAT(tanggal, '%Y') = ?", Request.Kode_user, year).Scan(&arr_invent.Total_pemasukan)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("transaksi_inventory").Select("IFNULL(SUM(total_harga),0) as total_pengeluaran").Where("kode_user = ? && DATE_FORMAT(tanggal, '%Y') = ? && status = 0", Request.Kode_user, year).Scan(&arr_invent.Total_pembayaran_pending)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	for i := 1; i <= 12; i++ {
		var chart_pengeluaran response.Chart_Pengeluaran
		var chart_pemasukan response.Chart_Pemasukan
		year_month := ""
		month_year := ""

		if i < 10 {
			year_month = year + "-" + "0" + strconv.Itoa(i)
			month_year = "0" + strconv.Itoa(i) + "-" + year
		} else {
			year_month = year + "-" + strconv.Itoa(i)

			month_year = strconv.Itoa(i) + "-" + year
		}

		err = con.Table("transaksi_inventory").Select("DATE_FORMAT(tanggal, '%b') AS tanggal", "IFNULL(SUM(total_harga),0) as total_pengeluaran", "DATE_FORMAT(tanggal, '%Y-%m')").Where("kode_user = ? && DATE_FORMAT(tanggal, '%Y-%m') = ? && status = 1", Request.Kode_user, year_month).Group("DATE_FORMAT(tanggal, '%Y-%m')").Scan(&chart_pengeluaran)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		chart_pengeluaran.Tanggal = month_year

		err = con.Table("transaksi").Select("DATE_FORMAT(tanggal, '%b') AS tanggal", "IFNULL(SUM(total_harga),0) as total_pemasukan", "DATE_FORMAT(tanggal, '%Y-%m')").Where("kode_user = ? && DATE_FORMAT(tanggal, '%Y-%m') = ?", Request.Kode_user, year_month).Group("DATE_FORMAT(tanggal, '%Y-%m')").Scan(&chart_pemasukan)

		if err.Error == nil {
			res.Status = http.StatusNotFound
			res.Message = "Not Found"
			res.Data = arr_invent
		}

		chart_pemasukan.Tanggal = month_year

		arr_invent.Chart_Pemasukan = append(arr_invent.Chart_Pemasukan, chart_pemasukan)
		arr_invent.Chart_Pengeluaran = append(arr_invent.Chart_Pengeluaran, chart_pengeluaran)

	}

	if err.Error != nil {
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
