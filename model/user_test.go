package model

import (
	"fmt"
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	users, err := FindAllUsers()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if len(*users) == 0 {
		t.Fatalf("failed test %#v", err)
	}

	// for _, u := range users {
	// 	fmt.Println("name: " + u.Name)
	// }
}
func TestFindUserByID(t *testing.T) {
	user, err := FindUserByID(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if user.ID != 1 {
		t.Fatalf("failed test %#v", err)
	}
	if user.Name != "unokun" {
		t.Fatalf("failed test %#v", err)
	}
	//fmt.Println("name: " + user.Name)
}
func TestFindUserByEmail(t *testing.T) {
	user, err := FindUserByEmail("hoge@test.jp")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if user.ID != 10 {
		t.Fatalf("failed test %#v", err)
	}
	//fmt.Println("name: " + user.Name)
}
func TestFindUserByEmailAndPassword(t *testing.T) {
	err := FindUserByEmailAndPassword("hoge@test.jp", "hoge1234")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestCreateUser(t *testing.T) {
	user := User{
		Name:     "hoge3",
		Email:    "hoge3@test.jp",
		Password: "hoge1234",
		AreaID:   1,
		TrashFlg: "1",
	}

	// ユーザー登録
	err := CreateUser(&user)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	// 登録したユーザーの検索
	created, err := FindUserByEmail(user.Email)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	fmt.Println("id = " + string(created.ID))
	fmt.Println("name = " + created.Name)

	//登録したユーザーの削除
	_, err = DeleteUser(created.ID)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	err := UpdateAccessToken(13, "WgNmT9QwcnIohbHTDwGDhNr0AQrY3IAtmErzSLXhyxW")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
