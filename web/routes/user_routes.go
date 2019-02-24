package routes

import (
	"gc_alert/web/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unokun/gc_alert/model"
)

/*
サインアップ
*/
func UserSignUp(ctx *gin.Context) {
	println("post/signup")
	username := ctx.PostForm("user_name")
	email := ctx.PostForm("user_email")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("password_confirm")

	if password != passwordConf {
		println("Error: password and password_confirm not match")
		ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
		return
	}

	// ユーザー登録
	user := &model.User{
		Name:     username,
		Email:    email,
		Password: password,
	}

	err := model.CreateUser(user)
	if err != nil {
		log.Fatal(err)
		ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
		return
	}

	log.Println("Signup success!!")
	log.Println("  username: " + username)
	log.Println("  email: " + email)
	log.Println("  password: " + password)

	created, err := model.FindUserByEmail(user.Email)
	if err != nil {
		log.Fatal(err)
		ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
		return
	}

	session := sessions.GetDefaultSession(ctx)
	session.Set("user", created)
	session.Save()
	println("Session saved.")
	println("  sessionID: " + session.ID)
	ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
}

/*
サインイン
*/
func UserSignIn(ctx *gin.Context) {
	println("post/signin")
	email := ctx.PostForm("user_email")
	//password := ctx.PostForm("password")

	user, err := model.FindUserByEmail(email)
	if err != nil {
		println("Error: " + err.Error())
		ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
		return
	}

	println("Authentication Success!!")
	println("  username: " + user.Name)
	println("  email: " + user.Email)
	println("  password: " + user.Password)
	session := sessions.GetDefaultSession(ctx)
	session.Set("user", user)
	session.Save()
	model.UserAuthenticate(user)

	println("Session saved.")
	println("  sessionID: " + session.ID)
	ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
}

/*
ログアウト
*/
func UserLogOut(ctx *gin.Context) {
	session := sessions.GetDefaultSession(ctx)
	session.Terminate()
	ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
}
