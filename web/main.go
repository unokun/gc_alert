package main

import (
	"gc_alert/web/routes"
	"gc_alert/web/sessions"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/gc_alert/assets", "./assets")

	store := sessions.NewStore()
	router.Use(sessions.StartDefaultSession(store))

	user := router.Group("/gc_alert/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/signin", routes.UserSignIn)
		user.POST("/register_trashnotify", routes.UserRegisterTrashNotify)
	}

	user = router.Group("/gc_alert/callback")
	{
		user.POST("/authorize", routes.UserLineAuthorizeCallback)
		user.POST("/token", routes.UserLineTokenCallback)
	}

	router.GET("/gc_alert/", routes.Home)
	router.GET("/gc_alert/signin", routes.SignIn)
	router.GET("/gc_alert/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	router.Run(":9010")
}
