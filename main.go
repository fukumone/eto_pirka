package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.Static("/assets", "./assets")

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{"root": "password",}))

	authorized.GET("/", AdminHandler())

	router.GET("/", RootHandler())
	router.Run(":3000")
}
