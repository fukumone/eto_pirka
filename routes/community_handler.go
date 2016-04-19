package routes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/t-fukui/eto_pirka/models"
)

func CommunityShowHandler(c *gin.Context) {
	token.CreateToken()
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
		"Token": token.Id,
	})
}

func CommunityNewHandler(c *gin.Context) {
	token.CreateToken()
	router.LoadHTMLFiles("templates/layout.html", "templates/main/community/new.html")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"UserData": UserData,
		"Token": token.Id,
	})
}

// TODO:Flash Successメッセージ追加
func CommunityCreateHandler(c *gin.Context) {
	var form models.CommunityForm
	c.Bind(&form)
	userId, _ := UserData["userid"].(string)
	community := models.Community{Name: form.Name, Description: form.Description, AdministratorId: userId}

	form.Community = community

	if models.ValidCommunity(&form) && form.Token == token.Id  {
		dbConnect.Debug().Create(&community)
		url := fmt.Sprintf("/user/%s", UserData["name"])
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		flashErrorMessage := FlashErrorMessage(c, store, "データを作成できませんでした")
		router.LoadHTMLFiles("templates/layout.html", "templates/main/community/new.html")
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"UserData": UserData,
			"Token": token.Id,
			"FlashErrorMessage": flashErrorMessage,
			"Errors": form.Errors,
		})
	}
}
