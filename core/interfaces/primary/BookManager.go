package primary

import "treinamento/core/domain"

type BookManager interface {
	CreateBook(book domain.Book) (int, error)
}
