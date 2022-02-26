package rest

import "github.com/gabrielmirandaBR/authenticate_go/internal/core/services"


type Book struct {
	bookService services.BookServices
}

func MusNewBook(service services.BookServices) *Book {
	return &Book{
		bookService: service,
	}
}