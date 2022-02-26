package main

import (
	"fmt"

	"github.com/gabrielmirandaBR/authenticate_go/adapters"
	"github.com/gabrielmirandaBR/authenticate_go/internal/core/rest"
	"github.com/gabrielmirandaBR/authenticate_go/internal/core/services"
	"github.com/gabrielmirandaBR/authenticate_go/internal/repository"
)

func main() {
	env, err := adapters.GetEnvironments()
	if err != nil {
		fmt.Printf("Error getting environments: %s", err)
		return
	}

	db := adapters.NewDatabase(&adapters.Config{
		Host: env.Host,
		Username: env.Username,
		Password: env.Password,
		DBName: env.DBName,
		Port: env.Port,
	})
	
	bookRepo := repository.MustNewBook(db)
	bookService := services.MustNewBook(bookRepo)
	bookRest := rest.MusNewBook(bookService)

	server := MustNewServer()
	router := server.Listen()
	// router.POST("/books", bookRest.)

	router.Run()
}
