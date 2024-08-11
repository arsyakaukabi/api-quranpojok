package main

import (
	"be-quranpojok/database"
	handler "be-quranpojok/internal/delivery/http" // Alias for the http package in your project
	"be-quranpojok/internal/repository"
	"be-quranpojok/internal/usecase"
	"log"
	nethttp "net/http" // Renamed import to avoid conflict

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.DBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ayahRepo := repository.NewAyahRepository(db)
	ayahUsecase := usecase.NewAyahUsecase(ayahRepo)
	ayahHandler := handler.NewAyahHandler(ayahUsecase) // Use the aliased package name

	r := mux.NewRouter()
	r.HandleFunc("/page/{page-number}", ayahHandler.GetAyahsByPage).Methods("POST")
	r.HandleFunc("/chapter-info/{surah_id}", ayahHandler.GetSurahInfo).Methods("GET")
	r.HandleFunc("/chapter/{surah_id}", ayahHandler.GetAyahsBySurah).Methods("POST")

	log.Println("Server started at :8080")
	nethttp.ListenAndServe(":8080", r) // Use the aliased nethttp
}
