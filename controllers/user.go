package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-1/models"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.Login(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func User_Profile(c echo.Context) error {
	kode_user := c.FormValue("kode_user")

	result, err := models.User_Profile(kode_user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
