package rest

import (
	"log"
	"net/http"

	"github.com/gabrielmirandaBR/authenticate_go/driver"
	"github.com/gabrielmirandaBR/authenticate_go/internal/core/domain"
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
	body := &driver.Book{}

	if err := ctx.ShouldBindJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if body.Name == "" || body.Author == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "name and author are requireds",
		})

		return
	}

	book := &domain.Book{
		Name:        body.Name,
		Description: body.Description,
		MediumPrice: body.MediumPrice,
		Author:      body.Author,
		ImageURL:    body.ImageURL,
	}

	newBook, err := b.bookService.CreateBook(ctx, book)
	if err != nil {
		log.Fatalf("%s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"ID":     newBook.ID,
		"Name":   newBook.Name,
		"Author": newBook.Author,
	})
}

func (b Book) DeleteBook(ctx *gin.Context) {

}
