package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/signature"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.Static("/assets", "./assets")

	// setup gomniauth
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		facebook.New(os.Getenv("FB_CLIENT_ID"), os.Getenv("FB_SECRET_KEY"), "http://localhost:3000/auth/callback/facebook"),
	)

	router.GET("/auth/login/facebook", LoginHandler())
	router.GET("/auth/callback/facebook", CallBackHandler())
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{os.Getenv("BasicAuthUSER"): os.Getenv("BasicAuthPASSWORD"),}))
	authorized.GET("/", AdminHandler())
	router.GET("/", RootHandler())
	router.Run(":3000")
}
