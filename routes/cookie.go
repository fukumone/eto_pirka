package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
)

func CookieSetup(c *gin.Context) {
	if authCookie, err := c.Request.Cookie("auth"); err == nil {
		if authCookie.Value != "" {
			UserData["UserData"] = objx.MustFromBase64(authCookie.Value)
		}
	}
}
