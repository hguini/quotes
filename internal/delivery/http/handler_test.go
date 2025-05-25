package http

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"quotes/internal/domain"
	"quotes/internal/infrastructure/memory"
	"quotes/internal/usecase"
	"strconv"
	"testing"
)

func setupTestServer() http.Handler {
	repo := memory.NewInMemoryRepo()
	uc := usecase.NewQuoteUsecase(repo)
	handler := NewQuoteHandler(uc)
	return NewRouter(handler)
}

func TestCreateAndGetQuotes(t *testing.T) {
	router := setupTestServer()

	body := `{"author":"Confucius","quote":"Life is simple"}`
	req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	require.Equal(t, http.StatusCreated, res.Code)

	var created domain.Quote
	err := json.NewDecoder(res.Body).Decode(&created)
	require.NoError(t, err)
	require.Equal(t, "Confucius", created.Author)
	require.Equal(t, "Life is simple", created.Text)

	// GET all
	req = httptest.NewRequest(http.MethodGet, "/quotes", nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	require.Equal(t, http.StatusOK, res.Code)

	var all []domain.Quote
	err = json.NewDecoder(res.Body).Decode(&all)
	require.NoError(t, err)
	require.Len(t, all, 1)
}

func TestFilterByAuthor(t *testing.T) {
	router := setupTestServer()

	quotes := []domain.Quote{
		{Author: "A", Text: "Quote A"},
		{Author: "B", Text: "Quote B"},
	}

	for _, q := range quotes {
		b, _ := json.Marshal(q)
		req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
	}

	req := httptest.NewRequest(http.MethodGet, "/quotes?author=B", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	require.Equal(t, http.StatusOK, res.Code)

	var filtered []domain.Quote
	err := json.NewDecoder(res.Body).Decode(&filtered)
	require.NoError(t, err)
	require.Len(t, filtered, 1)
	require.Equal(t, "B", filtered[0].Author)
}

func TestGetRandomQuote(t *testing.T) {
	router := setupTestServer()

	// Добавим цитату
	quote := domain.Quote{Author: "X", Text: "Random test"}
	b, _ := json.Marshal(quote)
	req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	// Получим случайную
	req = httptest.NewRequest(http.MethodGet, "/quotes/random", nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	require.Equal(t, http.StatusOK, res.Code)

	var q domain.Quote
	err := json.NewDecoder(res.Body).Decode(&q)
	require.NoError(t, err)
	require.Equal(t, "X", q.Author)
}

func TestDeleteQuote(t *testing.T) {
	router := setupTestServer()

	// Создать цитату
	q := domain.Quote{Author: "ToDelete", Text: "To delete"}
	b, _ := json.Marshal(q)
	req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var created domain.Quote
	_ = json.NewDecoder(res.Body).Decode(&created)

	// Удалить её
	req = httptest.NewRequest(http.MethodDelete, "/quotes/"+strconv.Itoa(created.ID), nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	require.Equal(t, http.StatusNoContent, res.Code)

	// Повторное удаление — ошибка
	req = httptest.NewRequest(http.MethodDelete, "/quotes/"+strconv.Itoa(created.ID), nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)
	require.Equal(t, http.StatusNotFound, res.Code)
}

func TestInvalidCreate(t *testing.T) {
	router := setupTestServer()

	req := httptest.NewRequest(http.MethodPost, "/quotes", bytes.NewBufferString("invalid"))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	require.Equal(t, http.StatusBadRequest, res.Code)

	body, _ := io.ReadAll(res.Body)
	require.Contains(t, string(body), "invalid")
}
