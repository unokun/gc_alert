package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
)

func TestRequestGetAccessToken(t *testing.T) {
	err := requestGetAccessToken("eq8MhR9x3rvF8vZkypqgn5")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

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
