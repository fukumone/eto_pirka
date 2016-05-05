package routes

import (
	"github.com/gin-gonic/gin"
)

func FlashErrorMessage(c *gin.Context, message string) string {
	session, _ := store.Get(c.Request, "flash-message")

	session.AddFlash(message)
	flash := session.Flashes()
	return flash[0].(string)
}

func FlashSuccessMessage(c *gin.Context, message string) {
	session, _ := store.Get(c.Request, "flash-message")
	session.Values["flash-success"] = message
	session.Save(c.Request, c.Writer)
}

func GetSuccessMessage(c *gin.Context) string {
	session, _ := store.Get(c.Request, "flash-message")
	message := session.Values["flash-success"].(string)
	session.Values["flash-success"] = ""
	session.Save(c.Request, c.Writer)
	return message
}
