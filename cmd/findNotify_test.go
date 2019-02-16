package main

import (
	"testing"
)

func TestFindNotifyMessage_OK_F2(t *testing.T) {
	result, err := FindNotifyMessage(5, 2)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(result) == 0 {
		t.Fatal("failed test")
	}
}

func TestFindNotifyMessage_OK_M0(t *testing.T) {
	result, err := FindNotifyMessage(1, 0)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(result) == 0 {
		t.Fatal("failed test")
	}
}
func TestFindNotifyMessage_NG(t *testing.T) {
	result, err := FindNotifyMessage(6, 2)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(result) != 0 {
		t.Fatal("failed test")
	}
}
