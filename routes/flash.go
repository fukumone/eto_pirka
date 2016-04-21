package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func FlashErrorMessage(c *gin.Context, s *sessions.CookieStore, message string) string {
	session, _ := s.Get(c.Request, "flash-message")

	session.AddFlash(message)
	flash := session.Flashes()
	return flash[0].(string)
}
