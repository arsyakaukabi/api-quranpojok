package entity

type Ayah struct {
	ID          int    `json:"id"`
	SurahID     int    `json:"surah_id"`
	VerseNumber int    `json:"verse_number"`
	Text        string `json:"text"` // Add this field to hold the selected mushaf text
	PageNumber  int    `json:"page_number"`
	JuzNumber   int    `json:"juz_number"`
}
