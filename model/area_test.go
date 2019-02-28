package model

import "testing"

func TestFindAllPrefs(t *testing.T) {
	prefs, err := findAllPrefs()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*prefs) == 0 {
		t.Fatalf("failed test %#v", err)
	}
	for _, p := range *prefs {
		println("pref_id: " + p.PrefID)
		println("pref_name: " + p.PrefName)
	}
}
