package model

import "testing"

func TestFindAllPrefs(t *testing.T) {
	areas, err := FindAllPrefs()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*areas) == 0 {
		t.Fatalf("failed test %#v", err)
	}
	for _, a := range *areas {
		println("pref: " + a.PrefID + " " + a.PrefName)
	}
}
func TestFindCities(t *testing.T) {
	areas, err := FindCities(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*areas) == 0 {
		t.Fatalf("failed test %#v", err)
	}
	for _, a := range *areas {
		println("city: " + a.CityID + " " + a.CityName)
	}
}
func TestFindTowns(t *testing.T) {
	areas, err := FindTowns(1101)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*areas) == 0 {
		t.Fatalf("failed test %#v", err)
	}
	for _, a := range *areas {
		println("town: " + a.TownID + " " + a.TownName)
	}
}
func TestFindAreaByZipCd(t *testing.T) {
	areas, err := FindAreaByZipCd("703-8256")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*areas) == 0 {
		t.Fatalf("failed test %#v", err)
	}
	for _, a := range *areas {
		println("pref: " + a.PrefID + " " + a.PrefName)
		println("city: " + a.CityID + " " + a.CityName)
		println("town: " + a.TownID + " " + a.TownName)
	}
}
