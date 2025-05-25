package usecase_test

import (
	"quotes/internal/domain"
	"quotes/internal/infrastructure/memory"
	"quotes/internal/usecase"
	"testing"
)

func setupTestUsecase() *usecase.QuoteUseCase {
	repo := memory.NewInMemoryRepo()
	return usecase.NewQuoteUsecase(repo)
}

func TestAddAndListQuote(t *testing.T) {
	uc := setupTestUsecase()

	q := domain.Quote{Author: "Confucius", Text: "Life is simple"}
	_, err := uc.AddQuote(q)
	if err != nil {
		t.Fatalf("AddQuote failed: %v", err)
	}

	quotes, _ := uc.ListQuotes(nil)
	if len(quotes) != 1 {
		t.Errorf("Expected 1 quote, got %d", len(quotes))
	}

	if quotes[0].Author != "Confucius" {
		t.Errorf("Unexpected author: %s", quotes[0].Author)
	}
}

func TestFilterByAuthor(t *testing.T) {
	uc := setupTestUsecase()
	uc.AddQuote(domain.Quote{Author: "A", Text: "Text A"})
	uc.AddQuote(domain.Quote{Author: "B", Text: "Text B"})

	author := "A"
	result, _ := uc.ListQuotes(&author)
	if len(result) != 1 {
		t.Errorf("Expected 1 quote for author A, got %d", len(result))
	}
}

func TestRandomQuote(t *testing.T) {
	uc := setupTestUsecase()
	_, err := uc.RandomQuote()
	if err == nil {
		t.Errorf("Expected error when no quotes present")
	}

	uc.AddQuote(domain.Quote{Author: "A", Text: "X"})
	quote, err := uc.RandomQuote()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if quote.Text != "X" {
		t.Errorf("Unexpected quote text: %s", quote.Text)
	}
}

func TestDeleteQuote(t *testing.T) {
	uc := setupTestUsecase()
	q, _ := uc.AddQuote(domain.Quote{Author: "A", Text: "X"})
	err := uc.DeleteQuote(q.ID)
	if err != nil {
		t.Errorf("Expected delete to succeed, got: %v", err)
	}

	// Try deleting again
	err = uc.DeleteQuote(q.ID)
	if err == nil {
		t.Errorf("Expected error when deleting nonexistent quote")
	}
}
