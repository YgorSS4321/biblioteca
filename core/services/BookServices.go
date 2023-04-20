package services

import (
	"log"
	"treinamento/core/domain"
	"treinamento/core/interfaces/primary"
	"treinamento/core/interfaces/repository"
)

var _ primary.BookManager = (*BookServices)(nil)

type BookServices struct {
	bookRepository repository.BookLoader
}

func (service BookServices) CreateBook(book domain.Book) (int, error) {
	bookID, err := service.bookRepository.InsertBook(book)
	if err != nil {
		log.Print(err)
		return -1, nil
	}

	return bookID, nil
}
