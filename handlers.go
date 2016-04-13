package main

import (
	"fmt"
	"log"
	"net/http"
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth"

	"github.com/t-fukui/eto_pirka/models"
)

import gomniauthcommon "github.com/stretchr/gomniauth/common"

func RootHandler(c *gin.Context) {
	data := map[string]interface{}{}
	if authCookie, err := c.Request.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	Communities := []models.Community{}
	db.Debug().Find(&Communities)
	router.LoadHTMLFiles("templates/layout.html", "templates/index.html")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"Communities": Communities,
		"UserData": data,
	})
}

func AdminHandler(c *gin.Context) {
	router.LoadHTMLFiles("templates/layout.html", "templates/admin/index.html")
	c.HTML(http.StatusOK, "layout.html", nil)
}

func LoginHandler(c *gin.Context) {
	router.LoadHTMLFiles("templates/layout.html", "templates/login.html")
	c.HTML(http.StatusOK, "layout.html", nil)
}

func AuthHandler(c *gin.Context) {
	provider, err := gomniauth.Provider("facebook")
	if err != nil {
		log.Fatalln("Error when trying to get provider", provider, "-", err)
	}
	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Fatalln("Error when trying to GetBeginAuthURL for", provider, "-", err)
	}

	c.Writer.Header()["Location"] = []string{loginUrl}
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}

type User struct {
	gomniauthcommon.User
	uniqueID string
}

func (u User) UniqueID() string {
	return u.uniqueID
}

func CallBackHandler(c *gin.Context) {
	provider, err := gomniauth.Provider("facebook")
	if err != nil {
		log.Fatalln("Error when trying to get provider", provider, "-", err)
	}

	// get the credentials
	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(c.Request.URL.RawQuery))
	if err != nil {
		log.Fatalln("Error when trying to complete auth for", provider, "-", err)
	}

	user, err := provider.GetUser(creds)
	if err != nil {
		log.Fatalln("Error when trying to get user from", provider, "-", err)
	}
	User := &User{User: user}

	m := md5.New()
	User.uniqueID = fmt.Sprintf("%x", m.Sum(nil))

	// save some data
	authCookieValue := objx.New(map[string]interface{}{
		"userid":     User.uniqueID,
		"name":       user.Name(),
	}).MustBase64()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/"})

	c.Writer.Header()["Location"] = []string{"/"}
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}

func LogoutHandler(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
	})
	c.Writer.Header()["Location"] = []string{"/login"}
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}
