package model

import "github.com/unokun/gc_alert/db"

/*
県情報
*/
type Pref struct {
	PrefID   string `json:"pref_id"`
	PrefName string `json:"pref_name"`
}

/**
 */
func (p *Pref) TableName() string {
	return "areas"
}

/**
県一覧を検索する
*/
func findAllPrefs() (*[]Pref, error) {
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	prefs := []Pref{}
	result := db.Exec("SELECT DISTINCT pref_id, pref_name FROM areas ORDER BY pref_id").Scan(&prefs)
	return &prefs, result.Error
}
