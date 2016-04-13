package main

import (
	"os"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/signature"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/t-fukui/eto_pirka/config"
)

var (
	router *gin.Engine
	db gorm.DB
)

func init() {
	db = config.Database()
}

func main() {
	router = gin.Default()
	router.Use(FilterHandler)
	router.Static("/assets", "./assets")

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		facebook.New(os.Getenv("FB_CLIENT_ID"), os.Getenv("FB_SECRET_KEY"), os.Getenv("FB_HOST")),
	)
	router.GET("/login", LoginHandler)
	router.GET("/auth/login/facebook", AuthHandler)
	router.GET("/auth/callback/facebook", CallBackHandler)
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{os.Getenv("BasicAuthUSER"): os.Getenv("BasicAuthPASSWORD"),}))
	authorized.GET("/", AdminHandler)
	router.GET("/", RootHandler)
	router.GET("/logout", LogoutHandler)
	router.Run(":3000")
}
