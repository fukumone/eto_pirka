package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/t-fukui/eto_pirka/models"
)

func RootHandler(c *gin.Context) {
	Communities := []models.Community{}
	dbConnect.Debug().Find(&Communities)
	router.LoadHTMLFiles("templates/layout.html", "templates/index.html")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Communities": Communities,
		"UserData":    UserData,
	})
}
