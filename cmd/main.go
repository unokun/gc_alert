package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type NotifyMessage struct {
	AccessToken string `json:"access_token"` // access_token
	Message     string `json:"message"`      // message
}
type NotifyMessages []*NotifyMessage

// DBConnect returns *sql.DB
func DBConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "agent"
	dbPass := os.Getenv("MYSQL_TRASH_PASSWORD") // 環境変数から取得
	dbName := "notifydb"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// sendLineNotifyはメッセージmsgをLINE Notifyに送信します。
func sendLineNotify(access_token string, msg string) error {
	fmt.Println(access_token)
	fmt.Println(msg)
	var url = "https://notify-api.line.me/api/notify"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("message="+msg)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+access_token)

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
func SendNotifyMessage(messages NotifyMessages) error {
	for _, m := range messages {
		err := sendLineNotify(m.AccessToken, m.Message)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
func FindNotifyMessage(dayOfWeek int, week int) (NotifyMessages, error) {
	db := DBConnect()
	result, err := db.Query("select T1.access_token, T2.message from users T1 inner join trash_notifies T2 on T2.area_id = T1.area_id where T2.day_of_week = ? and (T2.week = ? OR T2.week = 0) ORDER BY T1.area_id", dayOfWeek, week)
	if err != nil {
		panic(err.Error())
	}
	var messages NotifyMessages
	for result.Next() {
		var t string
		var m string

		err = result.Scan(&t, &m)
		if err != nil {
			panic(err.Error())
		}

		var message = new(NotifyMessage)
		message.AccessToken = t
		message.Message = m

		messages = append(messages, message)
	}
	return messages, err
}
func main() {
	// システム日時 -> 週/曜日
	var time = time.Now()
	var week = time.Day()/7 + 1
	var dayOfWeek = int(time.Weekday())

	// 通知
	messages, err := FindNotifyMessage(dayOfWeek, week)
	if err != nil {
		panic(err.Error())
	}

	err = SendNotifyMessage(messages)
	if err != nil {
		panic(err.Error())
	}
}
