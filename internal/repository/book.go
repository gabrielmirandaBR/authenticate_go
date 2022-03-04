package repository

import (
	"context"

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

func (b Book) CreateBook(ctx context.Context, book domain.Book)  error {
	
	return  nil
}

func (b Book) DeleteBook(ctx context.Context, book domain.Book)  error {
	
	return  nil
}
