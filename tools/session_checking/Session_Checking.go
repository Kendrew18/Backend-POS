package session_checking

import (
	"Bakend-POS/db"
	"Bakend-POS/models/response"
	"time"
)

func Session_Checking(UUID string) (response.User_Session_Response, bool) {
	var res response.User_Session_Response

	con := db.CreateConGorm()

	date := time.Now()
	tanggal := date.Format("2006-01-02")

	err := con.Table("user").Select("kode_user", "status").Where("uuid_session = ? AND date_session_invalid > ?", UUID, tanggal).Scan(&res)

	if err != nil {
		return res, false
	}

	if res.Kode_user == "" {
		return res, false
	} else {

		date := time.Now()
		tanggal_awal := date.Format("2006-01-02")
		date_akhir := date.AddDate(0, 0, 30)
		tanggal_terakhir := date_akhir.Format("2006-01-02")

		err = con.Table("user").Where("kode_user = ?", res.Kode_user).Update("date_last_open", tanggal_awal).Update("date_session_invalid", tanggal_terakhir)

		if err != nil {
			return res, false
		}

	}

	return res, true
}
