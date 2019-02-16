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

	store := sessions.NewDummyStore()
	router.Use(sessions.StartDefaultSession(store))

	user := router.Group("/gc_alert/user")
	{
		user.POST("/gc_alert/signup", routes.UserSignUp)
		user.POST("/gc_alert/signin", routes.UserSignIn)
	}

	router.GET("/gc_alert/", routes.Home)
	router.GET("/gc_alert/signin", routes.SignIn)
	router.GET("/gc_alert/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	router.Run(":9010")
}
