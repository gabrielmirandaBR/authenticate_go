package repository

import (
	"context"
	"log"

	"github.com/gabrielmirandaBR/authenticate_go/driven"
	"github.com/gabrielmirandaBR/authenticate_go/internal/core/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book domain.Book) error
	DeleteBook(ctx context.Context, book domain.Book) error
}

type Book struct {
	cli *gorm.DB
}

func MustNewBook(gormCli *gorm.DB) *Book {
	return &Book{
		cli: gormCli,
	}
}

func (b Book) CreateBook(ctx context.Context, book domain.Book) error {
	newBook := &driven.Book{
		ID:          book.ID,
		Name:        book.Name,
		Description: book.Description,
		MediumPrice: book.MediumPrice,
		Author:      book.Author,
		ImageURL:    book.ImageURL,
	}

	err := b.cli.Create(&newBook).Error
	if err != nil {
		log.Fatalf("%s", err)
		return err
	}

	return nil
}

func (b Book) DeleteBook(ctx context.Context, book domain.Book) error {

	return nil
}
