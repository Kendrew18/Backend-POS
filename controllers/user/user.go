package user

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/user"
	"Bakend-POS/tools/session_checking"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	var Request request.Login_Request

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := user.Login_User(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func SignUp(c echo.Context) error {
	var Request request.Sign_Up_Request

	err := c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := user.Sign_Up(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func UserProfile(c echo.Context) error {
	var Request request.Profile_User_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	//c.Response().Header().Set("X-Custom-Header", "Hello, World!")

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {

		result, err = user.User_Profile(Request)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

	} else {
		result.Status = http.StatusNotFound
		result.Message = "Session Invalid"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}

func UpdateUserProfile(c echo.Context) error {
	var Request request.Update_Profile_User_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {
		result, err = user.Update_User_Profile(Request)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

	} else {
		result.Status = http.StatusNotFound
		result.Message = "Session Invalid"
		result.Data = Request
	}

	return c.JSON(result.Status, result)
}
