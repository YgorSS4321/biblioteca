package request

import "treinamento/core/domain"

type Book struct {
	Title                 string `json:"title"`
	PagesQuantity         int    `json:"pages_quantity"`
	AuthorName            string `json:"author_name"`
	PublishingCompanyName string `json:"publishing_company_name"`
	PublicationYear       int    `json:"publication_year"`
}

func (dto Book) ToDomain() *domain.Book {
	return domain.NewBook(dto.Title, dto.PagesQuantity, dto.AuthorName, dto.PublishingCompanyName, dto.PublicationYear)
}
