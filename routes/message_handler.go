package routes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func MessageCreateHandler(c *gin.Context) {
	url := fmt.Sprintf("/community/%s", c.Params.ByName("id"))
	c.Redirect(http.StatusMovedPermanently, url)
}
