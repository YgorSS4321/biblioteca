package postgres

import (
	"log"
	"treinamento/core/domain"
	"treinamento/core/interfaces/repository"
)

var _ repository.BookLoader = (*BookRepository)(nil)

type BookRepository struct {
	Connector
}

func (repo BookRepository) InsertBook(book domain.Book) (int, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Print(err)
		return -1, err
	}
	defer repo.closeConnection(conn)

	query := `insert into book(title, pages_quantity, author_name, 
                 publishing_company_name, publication_year) values($1, $2, $3, $4, $5) 
                 returning id;`

	var bookID int
	err = conn.Get(
		&bookID,
		query,
		book.Title(),
		book.PagesQuantity(),
		book.AuthorName(),
		book.PublishingCompanyName(),
		book.PublicationYear(),
	)

	if err != nil {
		log.Print(err)
		return -1, err
	}

	return bookID, nil
}

func NewBookRepository(connectionManager Connector) *BookRepository {
	return &BookRepository{connectionManager}
}
