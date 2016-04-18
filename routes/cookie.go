package routes

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
)

func CookieSetup(c *gin.Context) {
	if authCookie, err := c.Request.Cookie("auth"); err == nil {
		if authCookie.Value != "" {
			UserData = objx.MustFromBase64(authCookie.Value)
			r := strings.NewReplacer(" ", "")
			str, _ := UserData["name"].(string)
			name := r.Replace(str)
			UserData["name"] = name
		}
	}
}
