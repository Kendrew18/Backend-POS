package news

import (
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"Bakend-POS/service/news"
	"Bakend-POS/tools/session_checking"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InputNews(c echo.Context) error {
	var Request request.Input_News_Request
	//var Request_session request.Token_Request
	var result response.Response
	//var err error

	data := c.FormValue("data")

	fmt.Println(data)

	jsonData := []byte(data)

	fmt.Println(jsonData)

	err := json.Unmarshal(jsonData, &Request)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Request)

	result, err = news.Input_News(Request, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadNewsUser(c echo.Context) error {
	var Request request.Read_News_Request
	var Request_session request.Token_Request
	var result response.Response
	var err error

	Request_session.Token = c.Request().Header.Get("token")

	User, condition := session_checking.Session_Checking(Request_session.Token)

	Request.Kode_user = User.Kode_user

	if condition {

		result, err = news.Read_News_User(Request)

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

func ReadNewsAdmin(c echo.Context) error {
	result, err := news.Read_News_Admin()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func DeleteNews(c echo.Context) error {
	var Request_body request.Delete_News_Request
	var err error

	err = c.Bind(&Request_body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := news.Delete_News(Request_body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
