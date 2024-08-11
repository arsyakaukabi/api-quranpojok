// ayah_usecase.go
package usecase

import (
	"be-quranpojok/internal/entity"
	"be-quranpojok/internal/repository"
)

type AyahUsecase interface {
	GetAyahsByPage(pageNumber int, mushaf string) ([]entity.Ayah, error)
	GetSurahInfo(surahID int) (entity.Surah, error)                    // Add this method
	GetAyahsBySurah(surahID int, mushaf string) ([]entity.Ayah, error) // Add this method
}

type ayahUsecase struct {
	ayahRepo repository.AyahRepository
}

func NewAyahUsecase(ayahRepo repository.AyahRepository) AyahUsecase {
	return &ayahUsecase{ayahRepo}
}

func (u *ayahUsecase) GetAyahsByPage(pageNumber int, mushaf string) ([]entity.Ayah, error) {
	return u.ayahRepo.GetAyahsByPage(pageNumber, mushaf)
}

func (u *ayahUsecase) GetSurahInfo(surahID int) (entity.Surah, error) {
	return u.ayahRepo.GetSurahInfo(surahID)
}

func (u *ayahUsecase) GetAyahsBySurah(surahID int, mushaf string) ([]entity.Ayah, error) {
	return u.ayahRepo.GetAyahsBySurah(surahID, mushaf)
}
