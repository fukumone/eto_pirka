package routes

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
	"github.com/gorilla/sessions"
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

func FlashErrorMessage(c *gin.Context, s *sessions.CookieStore, message string) string {
	session, _ := s.Get(c.Request, "flash-message")

	session.AddFlash(message)
	flash := session.Flashes()
	return flash[0].(string)
}
