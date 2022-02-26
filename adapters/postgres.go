package adapters

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(input *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		input.Host, 
		input.Username, 
		input.Password, 
		input.DBName, 
		input.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error creating database: %s", err)
		log.Fatal("error: ", err)
	}

	return db
}
