package routes

import (
	"gc_alert/web/config"
	"gc_alert/web/sessions"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(ctx *gin.Context) {
	var user *config.DummyUserModel

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unhappy home")
		println("  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
		return
	}

	user = buffer.(*config.DummyUserModel)
	println("Home sweet home")
	println("  sessionID: " + session.ID)
	println("  username: " + user.Username)
	println("  email: " + user.Email)

	session.Save()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"isSignIn": exists,
		"username": user.Username,
		"email":    user.Email,
	})
}

func SignIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signin.html", gin.H{})
}

func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
