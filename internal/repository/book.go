package repository

import (
	"gorm.io/gorm"
)

type BookRepository interface {
	
}

type Book struct {
	cli *gorm.DB
}

func MustNewBook(gormCli *gorm.DB) *Book {
	return &Book{
		cli: gormCli, 
	}
}
