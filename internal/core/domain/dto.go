package domain

import "time"

type Book struct {
	ID          string
	Name        string
	Description string
	MediumPrice float32
	Author      string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
