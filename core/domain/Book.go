package domain

type Book struct {
	title                 string
	pagesQuantity         int
	authorName            string
	publishingCompanyName string
	publicationYear       int
}

func (entity Book) Title() string {
	return entity.title
}

func (entity Book) PagesQuantity() int {
	return entity.pagesQuantity
}

func (entity Book) AuthorName() string {
	return entity.authorName
}

func (entity Book) PublishingCompanyName() string {
	return entity.publishingCompanyName
}

func (entity Book) PublicationYear() int {
	return entity.publicationYear
}

func NewBook(title string, pagesQuantity int, authorName, publishingCompanyName string, publicationYear int) *Book {
	return &Book{
		title:                 title,
		pagesQuantity:         pagesQuantity,
		authorName:            authorName,
		publishingCompanyName: publishingCompanyName,
		publicationYear:       publicationYear,
	}
}
