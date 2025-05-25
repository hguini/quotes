package usecase

import (
	"quotes/internal/domain"
)

type QuoteUseCase struct {
	repo domain.QuoteRepository
}

func NewQuoteUsecase(repo domain.QuoteRepository) *QuoteUseCase {
	return &QuoteUseCase{repo: repo}
}

func (uc *QuoteUseCase) AddQuote(q domain.Quote) (domain.Quote, error) {
	return uc.repo.Create(q)
}

func (uc *QuoteUseCase) ListQuotes(author *string) ([]domain.Quote, error) {
	if author != nil {
		return uc.repo.GetByAuthor(*author)
	}
	return uc.repo.GetAll()
}

func (uc *QuoteUseCase) RandomQuote() (domain.Quote, error) {
	return uc.repo.GetRandom()
}

func (uc *QuoteUseCase) DeleteQuote(id int) error {
	return uc.repo.Delete(id)
}
