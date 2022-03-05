package services

import (
	"context"
	"log"

	"github.com/gabrielmirandaBR/authenticate_go/internal/core/domain"
	"github.com/gabrielmirandaBR/authenticate_go/internal/repository"
	"github.com/google/uuid"
)

type BookServices interface {
	CreateBook(ctx context.Context, book *domain.Book) (domain.Book, error)
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

func (b Book) CreateBook(ctx context.Context, book *domain.Book) (domain.Book, error) {
	newID := uuid.New().String()

	newBook := domain.Book{
		ID: newID,
		Name: book.Name,
		Description: book.Description,
		MediumPrice: book.MediumPrice,
		Author: book.Author,
		ImageURL: book.ImageURL,
	}

	err := b.bookRepo.CreateBook(ctx, newBook)
	if err != nil {
		log.Fatalf("%s", err)
		return domain.Book{}, err
	}

	return newBook, nil
}

func (b Book) DeleteBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	
	return domain.Book{}, nil
}
