package main

import (
	"log"
	"net/http"
	delivery "quotes/internal/delivery/http"
	"quotes/internal/infrastructure/memory"
	"quotes/internal/usecase"
)

func main() {
	repo := memory.NewInMemoryRepo()
	uc := usecase.NewQuoteUsecase(repo)
	handler := delivery.NewQuoteHandler(uc)
	router := delivery.NewRouter(handler)

	log.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
