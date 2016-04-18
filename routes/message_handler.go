package routes

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/t-fukui/eto_pirka/models"
)

// TODO:Validation機能追加
func MessageCreateHandler(c *gin.Context) {
	community_id := c.Params.ByName("id")
	Community := models.Community{}
	dbConnect.Debug().First(&Community, community_id)

	Messages := []models.Message{}
	dbConnect.Debug().Where("community_id = ?", community_id).Find(&Messages)

	var form models.Message
	c.Bind(&form)
	CommunityId, _ := strconv.Atoi(c.Params.ByName("id"))
	message := models.Message{Name: form.Name, Body: form.Body, CommunityId: CommunityId}

	url := fmt.Sprintf("/community/%s", c.Params.ByName("id"))
	dbConnect.Debug().Create(&message)
	c.Redirect(http.StatusMovedPermanently, url)
}
