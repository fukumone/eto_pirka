package routes

import (
	"fmt"
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

func CommunityNewHandler(c *gin.Context) {
	router.LoadHTMLFiles("templates/layout.html", "templates/main/community/new.html")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"UserData": UserData,
	})
}

// TODO:Validation機能追加
func CommunityCreateHandler(c *gin.Context) {
	var form models.Community
	c.Bind(&form)
	community := models.Community{Name: form.Name, Description: form.Description}
	dbConnect.Debug().Create(&community)
	url := fmt.Sprintf("/user/%s", UserData["name"])
	c.Redirect(http.StatusMovedPermanently, url)
}
