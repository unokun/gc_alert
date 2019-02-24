package db

import (
	"testing"
)

func TestDBConnect(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer db.Close()
}
