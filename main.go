package main

import (
	"github.com/t-fukui/eto_pirka/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.RootHandler())
	router.Run(":3000")
}
