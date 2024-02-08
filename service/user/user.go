package user

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Login_User(Request request.Login_Request) (response.Response, error) {

	var res response.Response
	var us response.Login_Response
	con := db.CreateConGorm().Table("user")

	err := con.Select("kode_user", "status").Where("email =? AND password =?", Request.Email, Request.Password).Scan(&us).Error

	fmt.Println(err)

	if err != nil || us.Kode_user == "" {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		us.Kode_user = ""
		res.Data = us

	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = us
	}

	fmt.Println()

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

	err = con.Select("co", "kode_user", "nama_lengkap", "birth_date", "email", "category_bisnis", "nama_bisnis", "alamat_bisnis", "instagram", "facebook", "password").Create(&Request)

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

	con := db.CreateConGorm().Table("user")

	err := con.Select("kode_user", "nama_lengkap", "birth_date", "email", "category_bisnis", "nama_bisnis", "alamat_bisnis", "instagram", "facebook").Where("kode_user = ?", Request.Kode_user).Order("co ASC").Scan(&data).Error

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

	return res, nil
}
