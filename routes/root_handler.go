package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
	"github.com/t-fukui/eto_pirka/models"
)

func RootHandler(c *gin.Context) {
	data := map[string]interface{}{}
	if authCookie, err := c.Request.Cookie("auth"); err == nil || authCookie.Value == "" {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	Communities := []models.Community{}
	dbConnect.Debug().Find(&Communities)
	router.LoadHTMLFiles("templates/layout.html", "templates/index.html")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Communities": Communities,
		"UserData":    data,
	})
}
