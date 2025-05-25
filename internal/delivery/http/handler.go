package http

import (
	"encoding/json"
	"net/http"
	"quotes/internal/domain"
	"quotes/internal/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type QuoteHandler struct {
	uc *usecase.QuoteUseCase
}

func NewQuoteHandler(uc *usecase.QuoteUseCase) *QuoteHandler {
	return &QuoteHandler{uc: uc}
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var q domain.Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created, _ := h.uc.AddQuote(q)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var authorPtr *string
	if author != "" {
		authorPtr = &author
	}
	quotes, _ := h.uc.ListQuotes(authorPtr)
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandom(w http.ResponseWriter, r *http.Request) {
	quote, err := h.uc.RandomQuote()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(quote)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.uc.DeleteQuote(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
