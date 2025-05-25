package memory_test

import (
	"quotes/internal/domain"
	"quotes/internal/infrastructure/memory"
	"testing"
)

func TestInMemoryRepo(t *testing.T) {
	repo := memory.NewInMemoryRepo()

	// Test Create
	quote := domain.Quote{Author: "Test", Text: "Hello"}
	created, err := repo.Create(quote)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if created.ID != 1 {
		t.Errorf("Expected ID 1, got %d", created.ID)
	}

	// Test GetAll
	all, _ := repo.GetAll()
	if len(all) != 1 {
		t.Errorf("Expected 1 quote, got %d", len(all))
	}

	// Test GetByAuthor
	byAuthor, _ := repo.GetByAuthor("Test")
	if len(byAuthor) != 1 {
		t.Errorf("Expected 1 quote by author, got %d", len(byAuthor))
	}

	// Test GetRandom
	_, err = repo.GetRandom()
	if err != nil {
		t.Errorf("Expected random quote, got error: %v", err)
	}

	// Test Delete
	err = repo.Delete(created.ID)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	quotesAfterDelete, _ := repo.GetAll()
	if len(quotesAfterDelete) != 0 {
		t.Errorf("Expected 0 quotes after delete, got %d", len(quotesAfterDelete))
	}
}
