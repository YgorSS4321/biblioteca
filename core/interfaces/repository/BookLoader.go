package repository

import "treinamento/core/domain"

type BookLoader interface {
	InsertBook(book domain.Book) (int, error)
}
