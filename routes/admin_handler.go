package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	router.LoadHTMLFiles("templates/layout.html", "templates/admin/index.html")
	c.HTML(http.StatusOK, "layout.html", nil)
}
