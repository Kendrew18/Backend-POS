package transaksi

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/session_checking"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Input_Transaksi(Request request.Body_Input_Transaksi_Request) (response.Response, error) {
	var res response.Response

	User, condition := session_checking.Session_Checking(Request.Input_transaksi.Uuid_session)

	if condition {
		Request_barang := Request.Input_barang_transaksi
		Request := Request.Input_transaksi

		Request.Kode_user = User.Kode_user

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

		err = con.Table("transaksi").Select("co", "kode_transaksi", "kode_nota", "tanggal", "kode_jenis_pembayaran", "jumlah_total", "total_harga", "diskon", "kode_user").Create(&Request)

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

		err = con.Table("barang_transaksi").Select("co", "kode_barang_transaksi", "kode_transaksi", "kode_inventory", "jumlah_barang", "harga", "sub_total").Create(&Request_barang)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("transaksi").Where("kode_transaksi = ?", Request.Kode_transaksi).Update("total_harga", total_harga).Update("total_barang", total_jumlah)

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
	}

	return res, nil
}

func Read_Pembayaran() (response.Response, error) {
	var res response.Response

	return res, nil
}
