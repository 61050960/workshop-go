package api

import (
	controller "api/Controller"
	"fmt"

	// "github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)


func NewServer(port int) {
	r := echo.New() // new engine ไม่มี log recovery อยากใช้ค่อยใช้เอง
	// r.Use(gin.Recovery())
	r.GET("/", controller.Hello)
	r.GET("/users", controller.GetUser)
	r.Start(fmt.Sprintf(": %v", port))
}

