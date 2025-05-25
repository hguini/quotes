package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(h *QuoteHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/quotes", h.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", h.GetQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", h.GetRandom).Methods("GET")
	r.HandleFunc("/quotes/{id}", h.DeleteQuote).Methods("DELETE")
	return r
}
