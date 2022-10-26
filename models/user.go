package models

import (
	"fmt"
	"net/http"
	"project-1/db"
	str "project-1/struct"
)

func Login(username string, password string) (Response, error) {
	var res Response
	var us str.User

	con := db.CreateCon()

	sqlStatement := "SELECT kode_user, username, password FROM user where username=? && password=? "

	err := con.QueryRow(sqlStatement, username, password).Scan(&us.Kode_user, &us.Username, &us.Password)

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		us.Kode_user = ""
		res.Data = us
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = us

	fmt.Println(us)
	return res, nil
}

func User_Profile(kode_user string) (Response, error) {
	var res Response
	var us str.User_Profile

	con := db.CreateCon()

	sqlStatement := "SELECT nama_user FROM user where kode_user=?"

	err := con.QueryRow(sqlStatement, kode_user).Scan(&us.Nama_user)

	if err != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		us.Nama_user = ""
		res.Data = us
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = us

	fmt.Println(us)
	return res, nil
}
