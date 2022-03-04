package services

import (
	"context"

	"github.com/gabrielmirandaBR/authenticate_go/internal/core/domain"
	"github.com/gabrielmirandaBR/authenticate_go/internal/repository"
)

type BookServices interface {
	CreateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	DeleteBook(ctx context.Context, book domain.Book) (domain.Book, error)
}

type Book struct {
	bookRepo repository.BookRepository
}

func MustNewBook(repo repository.BookRepository) *Book {
	return &Book{
		bookRepo: repo,
	}
}

func (b Book) CreateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	
	return domain.Book{}, nil
}

func (b Book) DeleteBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	
	return domain.Book{}, nil
}
