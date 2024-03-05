package transaksi

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/session_checking"
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
		Request.Kode_pembayaran = "TR-" + strconv.Itoa(Request.Co)

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

		err = con.Table("transaksi").Select("co", "kode_pembayaran", "kode_nota", "tanggal", "kode_jenis_pembayaran", "kode_store", "jumlah_total", "total_harga", "kode_kasir", "diskon").Create(&Request)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}
	}

	return res, nil
}
