package session_checking

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/tools/decrypt"
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func Session_Checking(Token string) (response.User_Session_Response, bool) {
	var res response.User_Session_Response

	con := db.CreateConGorm()

	date := time.Now()
	tanggal := date.Format("2006-01-02")

	key_byte := []byte("S0ft1nd0PUtr4PErk4s4 s3cr3t K3ys")

	key := hex.EncodeToString(key_byte)

	fmt.Println(key)

	text := decrypt.Decrypt(key, Token)

	byteBuffer := bytes.NewBuffer([]byte(text))

	dec := gob.NewDecoder(byteBuffer)

	var token_data request.Token

	err2 := dec.Decode(&token_data)
	if err2 != nil {
		log.Fatal("decode error:", err2)
	}

	fmt.Println(token_data)

	err := con.Table("user").Select("kode_user", "status").Where(" date_session_invalid > ? AND kode_user = ?", tanggal, token_data.Kode_user).Scan(&res).Error

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

		err = con.Table("user").Where("kode_user = ?", res.Kode_user).Update("date_last_open", tanggal_awal).Update("date_session_invalid", tanggal_terakhir).Error

		if err != nil {
			return res, false
		}

	}

	return res, true
}
