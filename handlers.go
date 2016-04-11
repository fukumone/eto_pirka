package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/stretchr/objx"
	"github.com/stretchr/gomniauth"
)

func RootHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		router.LoadHTMLFiles("templates/layout.html", "templates/index.html")
		c.HTML(200, "layout.html", "index.html")
	}
}

func AdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		router.LoadHTMLFiles("templates/layout.html", "templates/admin/index.html")
		c.HTML(200, "layout.html", "admin/index.html")
	}
}

func LoginHandler() gin.HandlerFunc {
	provider, err := gomniauth.Provider("facebook")
	if err != nil {
		log.Fatalln("Error when trying to get provider", provider, "-", err)
	}
	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Fatalln("Error when trying to GetBeginAuthURL for", provider, "-", err)
	}

	return func(c *gin.Context) {
		c.Writer.Header()["Location"] = []string{loginUrl}
		c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func CallBackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// provider, err := gomniauth.Provider("facebook")
		// if err != nil {
		// 	log.Fatalln("Error when trying to get provider", provider, "-", err)
		// }

		// // get the credentials
		// creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		// if err != nil {
		// 	log.Fatalln("Error when trying to complete auth for", provider, "-", err)
		// }

		// user, err := provider.GetUser(creds)
		// if err != nil {
		// 	log.Fatalln("Error when trying to get user from", provider, "-", err)
		// }
		// chatUser := &chatUser{User: user}

		// m := md5.New()
		// io.WriteString(m, strings.ToLower(user.Email()))
		// chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))

		// avatarURL, err := avatars.GetAvatarURL(chatUser)
		// if err != nil {
		// 	log.Fatalln("Error when trying to GetAvatarURL", "-", err)
		// }

		// // save some data
		// authCookieValue := objx.New(map[string]interface{}{
		// 	"userid":     chatUser.uniqueID,
		// 	"name":       user.Name(),
		// 	"avatar_url": avatarURL,
		// }).MustBase64()

		// http.SetCookie(w, &http.Cookie{
		// 	Name:  "auth",
		// 	Value: authCookieValue,
		// 	Path:  "/"})

		c.Writer.Header()["Location"] = []string{"/"}
		c.Writer.WriteHeader(http.StatusTemporaryRedirect)
	}
}
