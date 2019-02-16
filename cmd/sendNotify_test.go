package main

import (
	"testing"
)

func TestSendNotifyMessage(t *testing.T) {
	result, err := FindNotifyMessage(1, 1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(result) != 1 {
		t.Fatal("failed test")
	}

	SendNotifyMessage(result)
	/*
			err = sendLineNotify("Oi69DSaj8Hdymi1VYFvh0Aqz3wdzMAMlS66wsS5Y6z6", "aaa")
		    if err != nil {
		        t.Fatalf("failed test %#v", err)
			}
	*/
}
