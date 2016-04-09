package handlers

import (
	"github.com/gin-gonic/gin"
)

func RootHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "index.tmpl", "")
	}
}
