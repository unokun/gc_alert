package routes

import (
	"bytes"
	"gc_alert/web/sessions"
	"strings"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unokun/gc_alert/model"
)

type ACCESS_TOKEN struct {
	AccessToken string `form:"access_token" json:"access_token" binding:"required"`
}

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

/*
 */
func UserRegisterTrashNotify(ctx *gin.Context) {

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unhappy home")
		println("  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
		return
	}

	var user *model.User
	user = buffer.(*model.User)

	url := createRequestAuthorizeURL(user.ID)
	println("redirect to: " + url)

	session.Save()

	ctx.Redirect(http.StatusSeeOther, url)

}

/*
Authorize Request URLを作成する。
*/
func createRequestAuthorizeURL(userID int) string {

	// CSRF 攻撃に対応するためのsessionIDを元にトークンを作成
	//h := sha1.New()
	//hash := hex.EncodeToString(h.Sum([]byte(sessionID)))

	var builder strings.Builder
	builder.WriteString("https://notify-bot.line.me/oauth/authorize?")
	builder.WriteString("response_type=code&")
	builder.WriteString("client_id=fmvHNOiimeuehStxOKXsVA&")
	builder.WriteString("redirect_uri=https://smaphonia.jp/gc_alert/callback/authorize&")
	builder.WriteString("scope=notify&")
	builder.WriteString("state=" + string(userID) + "&")
	builder.WriteString("response_mode=form_post")
	return builder.String()

}

/*
Line Authorizeリクエストのコールバック
*/
func UserLineAuthorizeCallback(ctx *gin.Context) {
	code := ctx.PostForm("code")
	//state := ctx.PostForm("state")
	// [TODO]セッションIDからトークンを作成し改ざんされていないことを確認する
	session := sessions.GetDefaultSession(ctx)
	_, exists := session.Get("user")
	if exists {
		println("sesseion user exists")
	} else {
		println("sesseion user not exitst")
	}

	println("code = " + code)

	requestGetAccessToken(code)

}

/*
 */
func UserLineTokenCallback(ctx *gin.Context) {
	session := sessions.GetDefaultSession(ctx)
	_, exists := session.Get("user")
	if exists {
		println("sesseion user exists")
	} else {
		println("sesseion user not exitst")
	}
	var json ACCESS_TOKEN
	if ctx.BindJSON(&json) == nil {
		println("access_token = " + json.AccessToken)

		// DB更新
	}
	/*
		session := sessions.GetDefaultSession(ctx)
		buffer, exists := session.Get("user")
		if !exists {
			println("Unhappy home")
			println("  sessionID: " + session.ID)
			session.Save()
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
			return
		}

		var user *model.User
		user = buffer.(*model.User)

		println("session id = " + session.ID)
		println("user id = " + string(user.ID))

		session.Save()

		var json ACCESS_TOKEN
		if ctx.BindJSON(&json) == nil {
			println("access_token = " + json.AccessToken)

			// DB更新
		}
	*/
}

/*
ACCESS_TOKEN取得をリクエストします
*/
func requestGetAccessToken(code string) error {
	var url = "https://notify-bot.line.me/oauth/token"

	var builder strings.Builder
	builder.WriteString("grant_type=authorization_code&")
	builder.WriteString("code=" + code + "&")
	builder.WriteString("redirect_uri=https://smaphonia.jp/gc_alert/callback/token&")
	builder.WriteString("client_id=fmvHNOiimeuehStxOKXsVA&")
	builder.WriteString("client_secret=XuMKCv7Y0zFxGvUmrkoj03h6GQuRt1m34fPOTun5EEC")
	content := builder.String()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(content)))
	if err != nil {
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()

	return err
}
