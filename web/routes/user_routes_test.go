package routes

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
)

func TestHttpPost(t *testing.T) {
	var apiurl = "https://notify-bot.line.me/oauth/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", "aaa")
	values.Add("redirect_uri", "https://smaphonia.jp/gc_alert/callback/token")
	values.Add("client_id", "fmvHNOiimeuehStxOKXsVA")
	values.Add("client_secret", "XuMKCv7Y0zFxGvUmrkoj03h6GQuRt1m34fPOTun5EEC")

	println(values.Encode())
	println(strings.NewReader(values.Encode()))

	req, err := http.NewRequest("POST", apiurl, strings.NewReader(values.Encode()))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	dump, err := httputil.DumpRequest(req, true)
	println(string(dump))
}

func TestParseAccessToken(t *testing.T) {
	body := `{"status":200,"message":"access_token is issued","access_token":"HVet4KBcnLCTDkrJh2wKLtfDpdgKLUVDqv3TbuilCWG"}`
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(body)))

	token := AccessToken{}
	err := decoder.Decode(&token)
	if err != nil {
		println("err: " + err.Error())
		log.Fatal(err)
	}
	if len(token.Token) == 0 {
		log.Fatal("Parse access token failed.")
	}
	//println("access_token: " + token.Token)
}
