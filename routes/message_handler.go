package routes

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/t-fukui/eto_pirka/models"
)

type MessageForm struct {
	models.Message
	Token string
}

// TODO:Flash Successメッセージ追加
func MessageCreateHandler(c *gin.Context) {
	community_id := c.Params.ByName("id")
	Community := models.Community{}
	dbConnect.Debug().First(&Community, community_id)

	Messages := []models.Message{}
	dbConnect.Debug().Where("community_id = ?", community_id).Find(&Messages)

	var form MessageForm
	c.Bind(&form)
	CommunityId, _ := strconv.Atoi(c.Params.ByName("id"))
	name, _ := UserData["name"].(string)
	user_id, _ := UserData["userid"].(string)
	message := models.Message{Name: name, Body: form.Body, CommunityId: CommunityId, UserId: user_id}

	if models.ValidMessage(message) && form.Token == token.Id {
		url := fmt.Sprintf("/user/%s/community/show/%s", name, c.Params.ByName("id"))
		dbConnect.Debug().Create(&message)
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		flashErrorMessage := FlashErrorMessage(c, store, "データを作成できませんでした")

		router.LoadHTMLFiles("templates/layout.html", "templates/main/community/show.html")
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Community": Community,
			"Messages": Messages,
			"UserData": UserData,
			"FlashErrorMessage": flashErrorMessage,
			"Token": token.Id,
		})
	}
}

type DeleteForm struct {
	MessageId int
}

// TODO:Flash Successメッセージ追加
func MessageDeleteHandler(c *gin.Context) {
	token.CreateToken()
	community_id := c.Params.ByName("id")
	Community := models.Community{}
	dbConnect.Debug().First(&Community, community_id)

	Messages := []models.Message{}
	dbConnect.Debug().Where("community_id = ?", community_id).Find(&Messages)

	var form DeleteForm
	c.Bind(&form)

	name, _ := UserData["name"].(string)
	user_id, _ := UserData["userid"].(string)
	Message := models.Message{}
	dbConnect.Debug().First(&Message, form.MessageId)

	if models.CommnunityValidAdmin(Community, user_id) {
		url := fmt.Sprintf("/user/%s/community/show/%s", name, c.Params.ByName("id"))
		dbConnect.Debug().Delete(&Message)
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		flashErrorMessage := FlashErrorMessage(c, store, "管理者権限がないので削除できません")

		router.LoadHTMLFiles("templates/layout.html", "templates/main/community/show.html")
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Community": Community,
			"Messages": Messages,
			"UserData": UserData,
			"Token": token.Id,
			"FlashErrorDeleteMessage": flashErrorMessage,
		})
	}
}
