package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/t-fukui/eto_pirka/models"
)

func CommunityShowHandler(c *gin.Context) {
	community_id := c.Params.ByName("id")
	Community := models.Community{}
	dbConnect.Debug().First(&Community, community_id)

	Messages := []models.Message{}
	dbConnect.Debug().Where("community_id = ?", community_id).Find(&Messages)

	router.LoadHTMLFiles("templates/layout.html", "templates/main/community/show.html")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Community": Community,
		"Messages": Messages,
		"UserData": UserData,
	})
}
