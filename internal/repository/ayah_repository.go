// ayah_repository.go
package repository

import (
	"be-quranpojok/internal/entity"
	"database/sql"
	"fmt"
)

type AyahRepository interface {
	GetAyahsByPage(pageNumber int, mushaf string) ([]entity.Ayah, error)
	GetSurahInfo(surahID int) (entity.Surah, error)                    // Add this method
	GetAyahsBySurah(surahID int, mushaf string) ([]entity.Ayah, error) // Add this method
}

type ayahRepository struct {
	db *sql.DB
}

func NewAyahRepository(db *sql.DB) AyahRepository {
	return &ayahRepository{db}
}

func (r *ayahRepository) GetAyahsByPage(pageNumber int, mushaf string) ([]entity.Ayah, error) {
	var ayahs []entity.Ayah
	query := fmt.Sprintf(`
		SELECT id, surah_id, verse_number, %s as text, page_number, juz_number
		FROM ayahs
		WHERE page_number = $1
		ORDER BY id
    `, mushaf)

	rows, err := r.db.Query(query, pageNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ayah entity.Ayah
		if err := rows.Scan(&ayah.ID, &ayah.SurahID, &ayah.VerseNumber, &ayah.Text, &ayah.PageNumber, &ayah.JuzNumber); err != nil {
			return nil, err
		}
		ayahs = append(ayahs, ayah)
	}

	return ayahs, nil
}

func (r *ayahRepository) GetSurahInfo(surahID int) (entity.Surah, error) {
	var surah entity.Surah
	query := `
        SELECT id, name, english_name, ayah_count, bismillah_pre
        FROM surahs
        WHERE id = $1
    `
	err := r.db.QueryRow(query, surahID).Scan(&surah.ID, &surah.Name, &surah.EnglishName, &surah.AyahCount, &surah.BismillahPre)
	if err != nil {
		return surah, err
	}
	return surah, nil
}

func (r *ayahRepository) GetAyahsBySurah(surahID int, mushaf string) ([]entity.Ayah, error) {
	var ayahs []entity.Ayah
	query := fmt.Sprintf(`
        SELECT id, surah_id, verse_number, %s as text, page_number, juz_number
        FROM ayahs
        WHERE surah_id = $1
        ORDER BY id
    `, mushaf)

	rows, err := r.db.Query(query, surahID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ayah entity.Ayah
		if err := rows.Scan(&ayah.ID, &ayah.SurahID, &ayah.VerseNumber, &ayah.Text, &ayah.PageNumber, &ayah.JuzNumber); err != nil {
			return nil, err
		}
		ayahs = append(ayahs, ayah)
	}

	return ayahs, nil
}
