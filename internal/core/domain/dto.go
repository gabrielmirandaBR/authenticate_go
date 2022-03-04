package domain

import (
	"time"

	"gorm.io/gorm"
)


type Book struct {
		ID          uint           
		Name        string         
		Description string         
		MediumPrice float32        
		Author      string        
		ImageURL    string         
		CreatedAt   time.Time      
		UpdatedAt   time.Time      
		DeletedAt   gorm.DeletedAt 
}
