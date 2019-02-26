package routes

import (
	"testing"
)

func TestRequestGetAccessToken(t *testing.T) {
	err := requestGetAccessToken("eq8MhR9x3rvF8vZkypqgn5")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
