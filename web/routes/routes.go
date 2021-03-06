package routes

import (
	"gc_alert/web/sessions"
	"log"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unokun/gc_alert/model"
)

/*
 */
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
	/*
		println("Home sweet home")
		println("  sessionID: " + session.ID)
		println("  username: " + user.Name)
		println("  email: " + user.Email)
	*/

	session.Save()
	lineNotifyURL := createRequestAuthorizeURL(session.ID)
	println("lineNotifyURL: " + lineNotifyURL)
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"isSignIn":      exists,
		"username":      user.Name,
		"email":         user.Email,
		"isTrashFlg":    user.TrashFlg == "1",
		"lineNotifyURL": lineNotifyURL,
	})
}

/*
Authorize Request URLを作成する。
*/
func createRequestAuthorizeURL(sessionID string) string {

	// CSRF 攻撃に対応するためのsessionIDを元にトークンを作成
	//h := sha1.New()
	//hash := hex.EncodeToString(h.Sum([]byte(sessionID)))

	var builder strings.Builder
	builder.WriteString("https://notify-bot.line.me/oauth/authorize?")
	builder.WriteString("response_type=code&")
	builder.WriteString("client_id=fmvHNOiimeuehStxOKXsVA&")
	builder.WriteString("redirect_uri=https://smaphonia.jp/gc_alert/callback/authorize&")
	builder.WriteString("scope=notify&")
	builder.WriteString("response_mode=form_post&")
	builder.WriteString("state=")
	return builder.String()
}

/*
 */
func SignIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signin.html", gin.H{})
}

/*
 */
func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

/*
 */
func SearchArea(ctx *gin.Context) {
	zipCd, result := ctx.GetQuery("zip_cd")
	if !result {
		ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
		return
	}
	//println("zip_cd: " + zipCd)
	areas, err := model.FindAreaByZipCd(zipCd)
	if err != nil {
		log.Fatal(err)
		ctx.Redirect(http.StatusSeeOther, "/gc_alert/")
		return
	}

	ctx.JSON(http.StatusOK, *areas)
}

/*
 */
func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
