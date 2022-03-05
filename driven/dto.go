package driven

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	MediumPrice float32
	Author      string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
