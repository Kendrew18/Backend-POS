package user

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/encrypt"
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func Login_User(Request request.Login_Request) (response.Response, error) {

	var res response.Response
	var us response.Login_Response
	var token request.Token
	con := db.CreateConGorm()

	err := con.Table("user").Select("kode_user").Where("username =? AND password =?", Request.Username, Request.Password).Scan(&token.Kode_user).Error

	if err != nil || token.Kode_user == "" {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = us

	} else {

		key_byte := []byte("S0ft1nd0PUtr4PErk4s4 s3cr3t K3ys")

		key := hex.EncodeToString(key_byte)

		fmt.Println(key)

		uuid := uuid.NewString()

		token.Uuid_session = uuid

		date := time.Now()
		tanggal_awal := date.Format("2006-01-02")
		date_akhir := date.AddDate(0, 0, 30)
		tanggal_terakhir := date_akhir.Format("2006-01-02")

		var network bytes.Buffer        // Stand-in for a network connection
		enc := gob.NewEncoder(&network) // Will write to network.

		err := enc.Encode(token)

		if err != nil {
			log.Fatal("encode error:", err)
		}

		cryptoText := encrypt.Encrypt(key, network.String())

		fmt.Println(cryptoText)

		if err != nil {
			log.Fatalf("write file err: %v", err.Error())
		}

		err = con.Table("user").Where("kode_user = ?", token.Kode_user).Update("token", cryptoText).Update("date_last_open", tanggal_awal).Update("date_session_invalid", tanggal_terakhir).Error

		if err != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err
		}

		err = con.Table("user").Select("token", "status").Where("username =? AND password =?", Request.Username, Request.Password).Scan(&us).Error

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

	username := ""

	con := db.CreateConGorm().Table("user")

	err := con.Select("username").Where("username = ?", Request.Username).Order("co ASC").Scan(&username).Error

	if username == "" {

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

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

		Request.Key = uuid.NewString()

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Username Telah ada"
		return res, err
	}

	return res, nil
}

func User_Profile(Request request.Profile_User_Request) (response.Response, error) {
	var res response.Response
	var data response.User_Profile_Response

	con := db.CreateConGorm().Table("user")

	err := con.Select("kode_user", "nama_lengkap", "birth_date", "gender", "category_bisnis", "nama_bisnis", "alamat_bisnis", "telepon_bisnis", "email_bisnis", "instagram", "facebook").Where("kode_user = ?", Request.Kode_user).Order("co ASC").Scan(&data).Error

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
