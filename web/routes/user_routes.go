package routes

import (
	"encoding/json"
	"gc_alert/web/sessions"
	"strconv"
	"strings"

	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/unokun/gc_alert/model"
)

/*
 */
type AccessToken struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"access_token"`
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
*/

/*
Line Authorizeリクエストのコールバック
*/
func UserLineAuthorizeCallback(ctx *gin.Context) {
	code := ctx.PostForm("code")
	println("code = " + code)
	state := ctx.PostForm("state")
	println("state = " + state)
	error := ctx.PostForm("error")
	println("error = " + error)
	errorDescription := ctx.PostForm("error_description")
	println("error_description = " + errorDescription)

	// [TODO]セッションIDからトークンを作成し改ざんされていないことを確認する
	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if exists {
		println("sesseion user exists")
	} else {
		println("sesseion user not exitst")
	}

	var user *model.User
	user = buffer.(*model.User)

	requestGetAccessToken(user.ID, code, state)

	session.Save()
	println("Session saved.")
	println("  sessionID: " + session.ID)
	ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
}

/*
ACCESS_TOKEN取得をリクエストします
*/
func requestGetAccessToken(userID int, code string, state string) error {
	var apiurl = "https://notify-bot.line.me/oauth/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", "https://smaphonia.jp/gc_alert/callback/authorize")
	values.Add("client_id", "fmvHNOiimeuehStxOKXsVA")
	values.Add("client_secret", "XuMKCv7Y0zFxGvUmrkoj03h6GQuRt1m34fPOTun5EEC")

	req, err := http.NewRequest("POST", apiurl, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	dump, err := httputil.DumpRequest(req, true)
	println(string(dump))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("err: " + err.Error())
		//log.Fatal(err)
		return err
	}

	defer resp.Body.Close()
	dump, err = httputil.DumpResponse(resp, true)
	println(string(dump))
	println("status: " + resp.Status)

	if resp.StatusCode == 200 {
		decoder := json.NewDecoder(resp.Body)

		token := AccessToken{}
		err := decoder.Decode(&token)
		if err != nil {
			println("err: " + err.Error())
			log.Fatal(err)
		}
		println("access_token: " + token.Token)

		// DB登録
		areaID, err := strconv.Atoi(state)
		if err != nil {
			println("err: " + err.Error())
			log.Fatal(err)
		}
		println("areaID: " + string(areaID))
		model.UpdateAccessTokenAndAreaID(userID, token.Token, areaID)
	}
	return err
}
