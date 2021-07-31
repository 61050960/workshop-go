package controller

import (
	"api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func GetUsers(c *gin.Context) {

// 	us := service.UserService{}

// 	c.JSON(http.StatusOK, us.GetUsers())
// }

// func Hello(c *gin.Context) {
// 	c.JSON(http.StatusOK, "Hello world")
// }

//echo

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world")
}

func GetUser(c echo.Context) error {
	
	us := service.UserService{}
	return c.JSON(http.StatusOK, us.GetUser())
}
