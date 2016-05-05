package routes

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/signature"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"

	"github.com/stretchr/objx"
)

import gomniauthcommon "github.com/stretchr/gomniauth/common"

func init() {
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		facebook.New(os.Getenv("FB_CLIENT_ID"), os.Getenv("FB_SECRET_KEY"), os.Getenv("FB_HOST")),
	)
}

func LoginHandler(c *gin.Context) {
	router.LoadHTMLFiles("templates/main/layout.html", "templates/main/login.html")
	flashMessage := GetSuccessMessage(c)
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"FlashMessage": flashMessage,
	})
}

func AuthenticateHandler(c *gin.Context) {
	provider, err := gomniauth.Provider("facebook")
	if err != nil {
		log.Fatalln("Error when trying to get provider", provider, "-", err)
	}
	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Fatalln("Error when trying to GetBeginAuthURL for", provider, "-", err)
	}
	FlashSuccessMessage(c, "ログインしました")
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
		"userid": User.uniqueID,
		"name":   user.Name(),
	}).MustBase64()

	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/"})
	name, _ := UserData["name"].(string)
	url := fmt.Sprintf("/user/%s", name)
	c.Writer.Header()["Location"] = []string{url}
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}

func LogoutHandler(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "auth",
		Value: "",
		Path:  "/",
	})
	FlashSuccessMessage(c, "ログアウトしました")
	c.Writer.Header()["Location"] = []string{"/login"}
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}
