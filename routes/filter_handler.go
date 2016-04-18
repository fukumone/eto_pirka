package routes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func FilterHandler(c *gin.Context) {
	if cookie, err := c.Request.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		// not authenticated
		c.Writer.Header()["Location"] = []string{"/login"}
		c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		if UserData["name"] == nil {
					url := fmt.Sprintf("/user/%s", UserData["name"])
			c.Writer.Header()["Location"] = []string{url}
			c.Writer.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			c.Writer.Header()["Location"] = []string{"/login"}
			c.Writer.WriteHeader(http.StatusTemporaryRedirect)
		}
	}
}
