package model

import (
	"fmt"

	"github.com/unokun/gc_alert/db"
	"golang.org/x/crypto/bcrypt"
)

/*
ユーザーデータ
*/
type User struct {
	ID            int    `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string
	AccessToken   string
	AreaID        int    `json:"area_id"`
	TrashFlg      string `json:"trash_flg"`
	authenticated bool
}

/*
IsValidUser func(*User) bool
*/
func IsValidUser(user *User) bool {
	return true
}

/*
FindAllUsers func() (*[]User, error)
全ユーザーを検索する
*/
func FindAllUsers() (*[]User, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	users := &[]User{}
	result := db.Find(users)

	return users, result.Error
}

/*
FindUserByID func(int) (*User, error)
ユーザーを検索する
*/
func FindUserByID(id int) (*User, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	user := &User{}
	result := db.First(user, "ID = ?", id)

	return user, result.Error
}

/*
FindUserByEmail func() (*User, error)
EMailが一致するユーザーを検索する
*/
func FindUserByEmail(email string) (*User, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	user := &User{}
	result := db.First(user, "email = ?", email)

	return user, result.Error
}

/*
FindUserByEmailAndPassword func(email string, password string) (error)
メールアドレス、パスワードが一致するユーザーを検索する
*/
func FindUserByEmailAndPassword(email string, password string) error {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	user, err := FindUserByEmail(email)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("hash = " + user.Password)

	return compareHashAndPassword(user.Password, password)
}

/*
CreateUser func(user *User) (error)
ユーザーを登録する
*/
func CreateUser(user *User) error {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	passwordEnc, err := encryptPassword(user.Password)
	if err != nil {
		panic(err.Error())
	}
	user.Password = passwordEnc

	result := db.Create(user)
	if err != nil {
		panic(err.Error())
	}

	return result.Error
}

/*
DeleteUser func(id int) (*User, error)
ユーザーを削除する
*/
func DeleteUser(id int) (*User, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	user := &User{}
	result := db.Delete(user, "ID = ?", id)

	return user, result.Error
}

/*
UpdateUser func(user *User) (error)
ユーザーを更新する
*/
func UpdateUser(user *User) error {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result := db.Update(user)

	return result.Error
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

/*
ユーザー認証
*/
func UserAuthenticate(user *User) {
	user.authenticated = true
}
