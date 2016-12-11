package routes

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/fukumone/eto_pirka/models"
)

func MessageCreateHandler(c *gin.Context) {
	community_id := c.Params.ByName("id")
	Community := models.Community{}
	dbConnect.Debug().First(&Community, community_id)

	Messages := []models.Message{}
	dbConnect.Debug().Where("community_id = ?", community_id).Find(&Messages)
	var form models.MessageForm
	c.Bind(&form)

	CommunityId, _ := strconv.Atoi(c.Params.ByName("id"))
	name, _ := UserData["name"].(string)
	user_id, _ := UserData["userid"].(string)
	message := models.Message{Name: name, Body: form.Body, CommunityId: CommunityId, UserId: user_id}

	form.Message = message

	if models.ValidMessage(&form, token.Id) {

		url := fmt.Sprintf("/user/%s/community/show/%s", name, c.Params.ByName("id"))
		dbConnect.Debug().Create(&message)
		FlashSuccessMessage(c, "メッセージの作成に成功しました")
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		flashErrorMessage := FlashErrorMessage(c, "データを作成できませんでした")
		router.LoadHTMLFiles("templates/main/layout.html", "templates/main/community/show.html")
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Community": Community,
			"Messages": Messages,
			"UserData": UserData,
			"FlashErrorMessage": flashErrorMessage,
			"Token": token.Id,
			"Errors": form.Errors,
		})
	}
}

type DeleteForm struct {
	MessageId int
}

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
		FlashSuccessMessage(c, "メッセージの削除に成功しました")
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		flashErrorMessage := FlashErrorMessage(c, "管理者権限がないので削除できません")

		router.LoadHTMLFiles("templates/main/layout.html", "templates/main/community/show.html")
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"Community": Community,
			"Messages": Messages,
			"UserData": UserData,
			"Token": token.Id,
			"FlashErrorDeleteMessage": flashErrorMessage,
		})
	}
}
