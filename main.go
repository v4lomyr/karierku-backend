package main

import (
	"fmt"
	"net/http"

	handlers "karierku.com/backend/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandlerLanding)
	http.HandleFunc("/quiz", handlers.HandlerQuiz)
	// http.HandleFunc("/result", handlers.HandlerResult)
	http.HandleFunc("/lowongan", handlers.HandlerLowongan)
	http.HandleFunc("/lowongan/reccomendation", handlers.HandlerRekomendasiLowongan)

	fmt.Println("Server Online")
	http.ListenAndServe(":9000", nil)
}