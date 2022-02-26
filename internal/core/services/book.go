package services

import "github.com/gabrielmirandaBR/authenticate_go/internal/repository"

type BookServices interface {

}

type Book struct {
	bookRepo repository.BookRepository
}

func MustNewBook(repo repository.BookRepository) *Book {
	return &Book{
		bookRepo: repo,
	}
}
