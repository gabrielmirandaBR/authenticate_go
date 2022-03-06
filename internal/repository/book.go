package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/gabrielmirandaBR/authenticate_go/driven"
	"github.com/gabrielmirandaBR/authenticate_go/internal/core/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book domain.Book) error
	DeleteBook(ctx context.Context, book domain.Book) error
	FindBook(ctx context.Context, bookID string)  error
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
		return fmt.Errorf("this book could not be created")
	}

	return nil
}

func (b Book) FindBook(ctx context.Context, bookID string)  error {
	book := &driven.Book{
		ID: bookID,
	}

	err := b.cli.First(&book).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("this book could not be found or has already been deleted")
	}

	return  nil
}

func (b Book) DeleteBook(ctx context.Context, book domain.Book) error {
	bookDeleted := &driven.Book {
		ID: book.ID,
	}

	err := b.cli.Where("id = ?", bookDeleted.ID).Delete(&bookDeleted).Error
	if err != nil {
		return fmt.Errorf("this book could not be deleted")
	}

	return nil
}
