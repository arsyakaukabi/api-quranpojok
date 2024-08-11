// ayah_handler.go
package http

import (
	"be-quranpojok/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AyahHandler struct {
	ayahUsecase usecase.AyahUsecase
}

func NewAyahHandler(ayahUsecase usecase.AyahUsecase) *AyahHandler {
	return &AyahHandler{ayahUsecase}
}

func (h *AyahHandler) GetAyahsByPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageNumber, err := strconv.Atoi(vars["page-number"])
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	var body struct {
		Mushaf string `json:"mushaf"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ayahs, err := h.ayahUsecase.GetAyahsByPage(pageNumber, body.Mushaf)
	if err != nil {
		http.Error(w, "Error fetching ayahs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ayahs)
}

func (h *AyahHandler) GetSurahInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	surahID, err := strconv.Atoi(vars["surah_id"])
	if err != nil {
		http.Error(w, "Invalid surah ID", http.StatusBadRequest)
		return
	}

	surah, err := h.ayahUsecase.GetSurahInfo(surahID)
	if err != nil {
		http.Error(w, "Error fetching surah info", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(surah)
}

func (h *AyahHandler) GetAyahsBySurah(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	surahID, err := strconv.Atoi(vars["surah_id"])
	if err != nil {
		http.Error(w, "Invalid surah ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Mushaf string `json:"mushaf"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ayahs, err := h.ayahUsecase.GetAyahsBySurah(surahID, body.Mushaf)
	if err != nil {
		http.Error(w, "Error fetching ayahs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ayahs)
}
