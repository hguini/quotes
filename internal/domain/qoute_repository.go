package domain

type QuoteRepository interface {
	Create(quote Quote) (Quote, error)
	GetAll() ([]Quote, error)
	GetByAuthor(author string) ([]Quote, error)
	GetRandom() (Quote, error)
	Delete(id int) error
}
