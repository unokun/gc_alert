package db

import (
	//	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
Connect func() (*gorm.DB, error)
DBに接続する
*/
func Connect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "agent"
	PASS := "agent"
	DBNAME := "notifydb"
	PROTOCOL := "tcp(localhost:3306)"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db, err
}
