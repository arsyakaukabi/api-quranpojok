// surah.go
package entity

type Surah struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	EnglishName  string `json:"english_name"`
	AyahCount    int    `json:"ayah_count"`
	BismillahPre bool   `json:"bismillah_pre"`
}
