package main

import (
	"github.com/gin-gonic/gin"
)

func RootHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		router.LoadHTMLFiles("templates/layout.html", "templates/index.html")
		c.HTML(200, "layout.html", "index.html")
	}
}

func AdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		router.LoadHTMLFiles("templates/layout.html", "templates/admin/index.html")
		c.HTML(200, "layout.html", "admin/index.html")
	}
}
