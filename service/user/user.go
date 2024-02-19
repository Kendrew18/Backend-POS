package user

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/session_checking"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func Login_User(Request request.Login_Request) (response.Response, error) {

	var res response.Response
	var us response.Login_Response
	con := db.CreateConGorm()

	kode_user := ""

	err := con.Table("user").Select("kode_user").Where("username =? AND password =?", Request.Username, Request.Password).Scan(&kode_user).Error

	if err != nil || kode_user == "" {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = us

	} else {
		uuid := uuid.NewString()

		date := time.Now()
		tanggal_awal := date.Format("2006-01-02")
		date_akhir := date.AddDate(0, 0, 30)
		tanggal_terakhir := date_akhir.Format("2006-01-02")

		err = con.Table("user").Where("kode_user = ?", kode_user).Update("uuid_session", uuid).Update("date_last_open", tanggal_awal).Update("date_session_invalid", tanggal_terakhir).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err
		}

		err = con.Table("user").Select("uuid_session", "status").Where("username =? AND password =?", Request.Username, Request.Password).Scan(&us).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err
		}

		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = us
	}

	return res, nil
}

func Sign_Up(Request request.Sign_Up_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm().Table("user")

	co := 0

	err := con.Select("co").Order("co DESC").Scan(&co)

	Request.Co = co + 1
	Request.Kode_user = "US-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Birth_date)
	Request.Birth_date = date.Format("2006-01-02")

	err = con.Select("co", "kode_user", "nama_lengkap", "birth_date", "gender", "category_bisnis", "nama_bisnis", "alamat_bisnis", "telepon_bisnis", "email_bisnis", "instagram", "facebook", "username", "password").Create(&Request)

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

func User_Profile(Request request.Profile_User_Request) (response.Response, error) {
	var res response.Response
	var data response.User_Profile_Response

	User, condition := session_checking.Session_Checking(Request.Uuid_session)

	if condition {
		con := db.CreateConGorm().Table("user")

		err := con.Select("kode_user", "nama_lengkap", "birth_date", "gender", "category_bisnis", "nama_bisnis", "alamat_bisnis", "telepon_bisnis", "email_bisnis", "instagram", "facebook").Where("kode_user = ?", User.Kode_user).Order("co ASC").Scan(&data).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err
		}

		if data.Kode_user == "" {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = data

		} else {
			res.Status = http.StatusOK
			res.Message = "Suksess"
			res.Data = data
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Session Invalid"
		res.Data = data
	}

	return res, nil
}

func Update_User_Profile(Request request.Update_Profile_User_Request) (response.Response, error) {
	var res response.Response
	con := db.CreateConGorm()

	date, _ := time.Parse("02-01-2006", Request.Birth_date)
	Request.Birth_date = date.Format("2006-01-02")

	err := con.Table("user").Where("kode_user = ?", Request.Kode_user).Select("nama_lengkap", "birth_date", "gender", "category_bisnis", "nama_bisnis", "alamat_bisnis", "telepon_bisnis", "email_bisnis", "instagram", "facebook").Updates(&Request)

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
