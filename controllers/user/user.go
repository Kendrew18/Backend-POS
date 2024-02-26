package user

import (
	"Bakend-POS/models/request"
	"Bakend-POS/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	var Request request.Login_Request

	err := c.Bind(&Request)

	if err != nil {
		panic(err)
	}

	result, err := user.Login_User(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func SignUp(c echo.Context) error {
	var Request request.Sign_Up_Request

	Request.Nama_lengkap = c.FormValue("nama_lengkap")
	Request.Birth_date = c.FormValue("birth_date")
	Request.Gender = c.FormValue("gender")
	Request.Category_bisnis = c.FormValue("category_bisnis")
	Request.Nama_bisnis = c.FormValue("nama_bisnis")
	Request.Alamat_bisnis = c.FormValue("alamat_bisnis")
	Request.Telepon_bisnis = c.FormValue("telepon_bisnis")
	Request.Email_bisnis = c.FormValue("email_bisnis")
	Request.Facebook = c.FormValue("facebook")
	Request.Instagram = c.FormValue("instagram")
	Request.Username = c.FormValue("username")
	Request.Password = c.FormValue("password")

	result, err := user.Sign_Up(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UserProfile(c echo.Context) error {
	var Request request.Profile_User_Request

	err := c.Bind(&Request)

	if err != nil {
		panic(err)
	}

	//Request.Uuid_session = c.FormValue("uuid_session")

	result, err := user.User_Profile(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UpdateUserProfile(c echo.Context) error {
	var Request request.Update_Profile_User_Request

	Request.Uuid_session = c.FormValue("uuid_session")
	Request.Nama_lengkap = c.FormValue("nama_lengkap")
	Request.Birth_date = c.FormValue("birth_date")
	Request.Gender = c.FormValue("gender")
	Request.Category_bisnis = c.FormValue("category_bisnis")
	Request.Nama_bisnis = c.FormValue("nama_bisnis")
	Request.Alamat_bisnis = c.FormValue("alamat_bisnis")
	Request.Telepon_bisnis = c.FormValue("telepon_bisnis")
	Request.Email_bisnis = c.FormValue("email_bisnis")
	Request.Facebook = c.FormValue("facebook")
	Request.Instagram = c.FormValue("instagram")

	result, err := user.Update_User_Profile(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
