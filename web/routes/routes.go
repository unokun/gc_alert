package routes

import (
	"gc_alert/web/sessions"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unokun/gc_alert/model"
)

func Home(ctx *gin.Context) {
	var user *model.User

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unhappy home")
		println("  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
		return
	}

	user = buffer.(*model.User)
	println("Home sweet home")
	println("  sessionID: " + session.ID)
	println("  username: " + user.Name)
	println("  email: " + user.Email)

	session.Save()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"isSignIn":   exists,
		"username":   user.Name,
		"email":      user.Email,
		"isTrashFlg": user.TrashFlg == "1",
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
