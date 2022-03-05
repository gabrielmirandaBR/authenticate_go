package main

import (
	"fmt"

	"github.com/gabrielmirandaBR/authenticate_go/adapters"
	"github.com/gabrielmirandaBR/authenticate_go/driven"
)

func main() {
	env, err := adapters.GetEnvironments()
	if err != nil {
		fmt.Printf("Error getting environments: %s", err)
		return
	}

	db := adapters.NewDatabase(&adapters.Config{
		Host:     env.Host,
		Username: env.Username,
		Password: env.Password,
		DBName:   env.DBName,
		Port:     env.Port,
	})

	tableBook := &driven.Book{}
	err = db.AutoMigrate(tableBook)
	if err != nil {
		panic(err)
	}
}
