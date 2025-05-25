package memory

import (
	"errors"
	"math/rand"
	"quotes/internal/domain"
	"sync"
)

type InMemoryRepo struct {
	mu     sync.RWMutex
	quotes []domain.Quote
	nextID int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{nextID: 1}
}

func (r *InMemoryRepo) Create(q domain.Quote) (domain.Quote, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	q.ID = r.nextID
	r.nextID++
	r.quotes = append(r.quotes, q)
	return q, nil
}

func (r *InMemoryRepo) GetAll() ([]domain.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.quotes, nil
}

func (r *InMemoryRepo) GetByAuthor(author string) ([]domain.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var res []domain.Quote
	for _, q := range r.quotes {
		if q.Author == author {
			res = append(res, q)
		}
	}
	return res, nil
}

func (r *InMemoryRepo) GetRandom() (domain.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if len(r.quotes) == 0 {
		return domain.Quote{}, errors.New("no quotes")
	}
	return r.quotes[rand.Intn(len(r.quotes))], nil
}

func (r *InMemoryRepo) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, q := range r.quotes {
		if q.ID == id {
			r.quotes = append(r.quotes[:i], r.quotes[i+1:]...)
			return nil
		}
	}
	return errors.New("quote not found")
}
