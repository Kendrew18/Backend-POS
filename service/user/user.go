package user

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/encrypt"
	"Bakend-POS/tools/googleverifid"
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func Random_Verif_Code(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func Send_Email(Email string, Username string, From string, App_name string) string {

	url := "https://send.api.mailtrap.io/api/send"
	method := "POST"

	var payload request.Payload
	var to request.To

	to.Email = Email

	verif_code := Random_Verif_Code(6)

	payload.From.Email = From + "@softindopp.com"
	payload.From.Name = "Softindo Putra Perkasa"
	payload.To = append(payload.To, to)
	payload.TemplateUUID = "be67ffe3-9749-4bff-ab1c-0b92a78d583a"
	payload.TemplateVariables.VerificationCode = verif_code
	payload.TemplateVariables.UserName = Username
	payload.TemplateVariables.AppsName = App_name

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "-1"
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))

	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	req.Header.Add("Authorization", "Bearer c3bf84bd500b2427337755537ce1914d")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "-1"
	}
	fmt.Println(string(body))

	return verif_code
}

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

func Sign_Up_With_Google(Request request.Sign_Up_Google) (response.Response, error) {
	var res response.Response
	var data response.Sign_Up_Google_Response

	token_info := googleverifid.VerifyGoogle(Request)

	var Request_Sign_UP request.Sign_Up_Request

	data.Email = token_info.Email
	data.Name = token_info.Name
	data.Client_id = token_info.Aud

	username := ""

	con := db.CreateConGorm().Table("user")

	err := con.Select("username").Where("username = ?", data.Name).Order("co ASC").Scan(&username).Error

	if username == "" {

		con := db.CreateConGorm().Table("user")

		co := 0

		err := con.Select("co").Order("co DESC").Limit(1).Scan(&co)

		Request_Sign_UP.Co = co + 1
		Request_Sign_UP.Kode_user = "US-" + strconv.Itoa(Request_Sign_UP.Co)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		Request_Sign_UP.Status = 0

		err = con.Select("co", "kode_user", "nama_lengkap", "birth_date", "gender", "category_bisnis", "nama_bisnis", "alamat_bisnis", "telepon_bisnis", "email_bisnis", "instagram", "facebook", "username", "password", "status").Create(&Request)

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

func Sign_Up(Request request.Sign_Up_Request) (response.Response, error) {
	var res response.Response

	username := ""

	con := db.CreateConGorm().Table("user")

	err := con.Select("username").Where("username = ? && email_bisnis = ? && status = 0", Request.Username, Request.Email_bisnis).Order("co ASC").Scan(&username).Error

	if username == "" {

		fmt.Println(Request)

		con := db.CreateConGorm().Table("user")

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
		Request.Status = -1

		Request.Key = uuid.NewString()

		err = con.Select("co", "kode_user", "nama_lengkap", "birth_date", "gender", "category_bisnis", "nama_bisnis", "alamat_bisnis", "telepon_bisnis", "email_bisnis", "instagram", "facebook", "username", "password", "status").Create(&Request)

		if err.Error != nil {
			fmt.Println("masuk")
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		var Request_OTP request.Otp_Request

		Request_OTP.Kode_otp = Send_Email(Request.Email_bisnis, Request.Nama_lengkap, "noreply", "Softipos")
		Request_OTP.Kode_user = Request.Kode_user

		Request_OTP.Email = Request.Email_bisnis

		layoutFormat := "2006-01-02 15:04:05"

		date_sent_time := time.Now()

		date_sent := date_sent_time.Format(layoutFormat)

		// date_resent_time, _ := time.Parse(layoutFormat, date_sent)

		// date_resent_time = date_resent_time.Add(time.Minute * time.Duration(1))

		// date_resent := date_resent_time.Format(layoutFormat)

		Request_OTP.Time_sent = date_sent
		Request_OTP.Nama_lengkap = Request.Nama_lengkap
		//Request_OTP.Time_resent = date_resent

		err = con.Select("kode_user", "nama_lengkap", "email", "kode_otp", "time_sent").Create(&Request_OTP)

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

func Resend_OTP(Request request.Resend_OTP_Request) (response.Response, error) {
	var res response.Response

	layoutFormat := "2006-01-02 15:04:05"

	con := db.CreateConGorm()

	time_sent := ""

	err := con.Table("otp").Select("time_sent", "nama_lengkap").Where("email = ?", Request.Email).Scan(&time_sent)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date_resent_time, _ := time.Parse(layoutFormat, time_sent)

	time_now := time.Now()

	dur := time_now.Sub(date_resent_time)

	if dur.Minutes() >= 1 {

		var Resent request.Update_OTP_Request

		Resent.Kode_otp = Send_Email(Request.Email, Request.Nama_lengkap, "noreply", "Softipos")

		layoutFormat := "2006-01-02 15:04:05"

		date_sent_time := time.Now()

		Resent.Time_sent = date_sent_time.Format(layoutFormat)

		err = con.Table("otp").Where("email =?", Request.Email).Select("time_sent", "kode_otp").Updates(&Resent)

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
		res.Message = "WAktu Belum 1 menit"
		return res, nil
	}

	return res, nil
}

func Activate_Account(Request request.Activate_Account_Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()

	kode_user := ""

	err := con.Table("otp").Select("kode_user").Where("email = ? && kode_otp = ?", Request.Email, Request.Kode_otp).Scan(&kode_user)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	if kode_user != "" {

		err = con.Table("otp").Where("email = ?", Request.Email).Delete("")

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		err = con.Table("user").Where("kode_user =?", kode_user).Update("status", 0)

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
		res.Message = "Kode OTP Salah"
		res.Data = Request
		return res, err.Error
	}

	return res, nil
}

// login with google
func Login_With_Google() {

}
