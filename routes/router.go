package routes

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"github.com/fukumone/eto_pirka/db"
)

var (
	router = gin.Default()
	dbConnect = db.InitDB()
	UserData = map[string]interface{}{}
	store = sessions.NewCookieStore([]byte("flash-message"))
)

func Init() *gin.Engine {
	router.Use(FilterHandler)
	router.Use(CookieSetup)
	router.Static("/assets", "./assets")

	// Admin(管理者の領域)
	// TODO コミュニティ、メッセージを強制削除機能追加
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{os.Getenv("BasicAuthUSER"): os.Getenv("BasicAuthPASSWORD"),}))
	authorized.GET("/", AdminHandler)

	// Main(Userの領域)
	router.GET("/login", LoginHandler)
	router.GET("/logout", LogoutHandler)
	router.GET("/auth/login/facebook", AuthenticateHandler)
	router.GET("/auth/callback/facebook", CallBackHandler)

	userRouter := router.Group("user/:name", func(c *gin.Context){
		if cookie, err := c.Request.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
			c.Abort()
		}
	})

	// Root Path
	userRouter.GET("/", RootHandler)

	community := userRouter.Group("/community")
	{
		community.GET("/show/:id", CommunityShowHandler)
		community.GET("/new", CommunityNewHandler)
		community.POST("/create", CommunityCreateHandler)
		community.POST("/show/:id/message/create", MessageCreateHandler)
		community.POST("/show/:id/message/delete", MessageDeleteHandler)
	}

	return router
}
