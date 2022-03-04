package rest

import (
	"github.com/gabrielmirandaBR/authenticate_go/internal/core/services"
	"github.com/gin-gonic/gin"
)


type Book struct {
	bookService services.BookServices
}

func MusNewBook(service services.BookServices) *Book {
	return &Book{
		bookService: service,
	}
}

func (b Book) CreateBook(ctx *gin.Context) {
	
}

func (b Book) DeleteBook(ctx *gin.Context) {
	
}
