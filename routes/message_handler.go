package routes

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/t-fukui/eto_pirka/models"
)

func MessageCreateHandler(c *gin.Context) {
	var form models.MessageForm
	c.Bind(&form)
	CommunityId, _ := strconv.Atoi(c.Params.ByName("id"))
	message := models.Message{Name: form.Name, Body: form.Body, CommunityId: CommunityId}
	url := fmt.Sprintf("/community/%s", c.Params.ByName("id"))
	dbConnect.Debug().Create(&message)
	c.Redirect(http.StatusMovedPermanently, url)
}
