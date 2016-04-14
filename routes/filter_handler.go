package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func FilterHandler(c *gin.Context) {
	if cookie, err := c.Request.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		// not authenticated
		c.Writer.Header()["Location"] = []string{"/login"}
		c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		// some other error
		panic(err.Error())
	} else {
		// success
		c.Writer.Header()["Location"] = []string{"/"}
		c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	}
}
