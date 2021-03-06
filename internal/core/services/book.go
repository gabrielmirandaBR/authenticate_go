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
	DeleteBook(ctx context.Context, book *domain.Book) error
	FindAllBooks(ctx context.Context) ([]domain.Book, error)
	UpdateBook(ctx context.Context, bookID string, book *domain.Book) error
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
		ID:          newID,
		Name:        book.Name,
		Description: book.Description,
		MediumPrice: book.MediumPrice,
		Author:      book.Author,
		ImageURL:    book.ImageURL,
	}

	if err := b.bookRepo.CreateBook(ctx, newBook); err != nil {
		log.Printf("%s", err)
		return domain.Book{}, err
	}

	return newBook, nil
}

func (b Book) findBook(ctx context.Context, bookID string) error {
	err := b.bookRepo.FindBook(ctx, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (b Book) DeleteBook(ctx context.Context, book *domain.Book) error {
	bookDeleted := domain.Book{
		ID: book.ID,
	}

	if err := b.findBook(ctx, bookDeleted.ID); err != nil {
		log.Printf("%s", err)
		return err
	}

	if err := b.bookRepo.DeleteBook(ctx, bookDeleted); err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}

func (b Book) FindAllBooks(ctx context.Context) ([]domain.Book, error) {
	books, err := b.bookRepo.FindAllBooks(ctx)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return books, nil
}

func (b Book) UpdateBook(ctx context.Context, bookID string, book *domain.Book) error {
	updateBook := domain.Book{
		Name:        book.Name,
		Description: book.Description,
		MediumPrice: book.MediumPrice,
		Author:      book.Author,
		ImageURL:    book.ImageURL,
	}

	err := b.bookRepo.UpdateBook(ctx, bookID, updateBook)
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}
