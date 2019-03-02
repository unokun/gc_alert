package model

import "github.com/unokun/gc_alert/db"

/*
県情報
*/
type Area struct {
	ID       string `json:"id"`
	PrefID   string `json:"pref_id"`
	PrefName string `json:"pref_name"`
	CityID   string `json:"city_id"`
	CityName string `json:"city_name"`
	TownID   string `json:"town_id"`
	TownName string `json:"town_name"`
}

/**
県一覧を検索する
*/
func FindAllPrefs() (*[]Area, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	prefs := []Area{}
	result := db.Raw("SELECT DISTINCT pref_id, pref_name FROM areas ORDER BY pref_id").Scan(&prefs)
	return &prefs, result.Error
}

/**
市一覧を検索する
*/
func FindCities(prefID int) (*[]Area, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	prefs := []Area{}
	result := db.Raw("SELECT DISTINCT city_id, city_name FROM areas WHERE pref_id = ? ORDER BY city_id", prefID).Scan(&prefs)
	return &prefs, result.Error
}

/**
町一覧を検索する
*/
func FindTowns(cityID int) (*[]Area, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	prefs := []Area{}
	result := db.Raw("SELECT DISTINCT town_id, town_name FROM areas WHERE city_id = ? ORDER BY town_id", cityID).Scan(&prefs)
	return &prefs, result.Error
}

/**
地域を検索する
*/
func FindAreaByZipCd(zipCD string) (*[]Area, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	prefs := []Area{}
	result := db.Where("zip_cd = ?", zipCD).Find(&prefs)
	return &prefs, result.Error
}
