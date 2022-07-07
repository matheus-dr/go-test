package entities

import "time"

type Book struct {
	Id        uint
	Title     string
	Author    Author
	Category  Category
	CreatedAt time.Time
	UpdatedAt time.Time
}
