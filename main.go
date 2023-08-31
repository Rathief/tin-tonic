package main

import (
	"tin-tonic/config"
	"tin-tonic/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	ah := handler.AuthHandler{DB: config.ConnectDB()}

	r := gin.Default()
	r.POST("/register", ah.Register)
	r.GET("/login", ah.Login)

	r.Run()
}
