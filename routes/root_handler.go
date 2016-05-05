package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/t-fukui/eto_pirka/models"
)

// TODO paginationを実装する(Community一覧)
func RootHandler(c *gin.Context) {
	Communities := []models.Community{}
	dbConnect.Debug().Find(&Communities)
	router.LoadHTMLFiles("templates/main/layout.html", "templates/main/index.html")
	flashSuccessMessage := GetSuccessMessage(c)
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Communities": Communities,
		"UserData":    UserData,
		"FlashSuccessMessage": flashSuccessMessage,
	})
}
