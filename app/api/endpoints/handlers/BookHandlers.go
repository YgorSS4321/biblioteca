package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"treinamento/app/api/endpoints/dto/request"
	"treinamento/app/api/endpoints/dto/response"
	"treinamento/core/interfaces/primary"
)

type BookHandlers struct {
	bookService primary.BookManager
}

func (handler BookHandlers) PostBook(c echo.Context) error {
	var bookDTO request.Book
	if err := c.Bind(&bookDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(http.StatusBadRequest,
			"Algo está incorreto em sua requisição"),
		)
	}

	book := bookDTO.ToDomain()
	bookID, err := handler.bookService.CreateBook(*book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError(http.StatusInternalServerError,
			"Parece que algo deu errado em sua base de dados"),
		)
	}

	return c.JSON(http.StatusCreated, &response.Created{ID: bookID})
}

func NewBookHandlers(bookService primary.BookManager) *BookHandlers {
	return &BookHandlers{bookService: bookService}
}
