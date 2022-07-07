package dto

import (
	"github.com/matheus-dr/go-test/src/entities"
	"time"
)

type CreateBookDto struct {
	Title      string `json:"title"`
	AuthorId   uint   `json:"authorId"`
	CategoryId uint   `json:"categoryId"`
}

type ListBookDto struct {
	Id       uint              `json:"id"`
	Title    string            `json:"title"`
	Author   entities.Author   `json:"author"`
	Category entities.Category `json:"category"`
}

type FindBookDto struct {
	Id        uint              `json:"id"`
	Title     string            `json:"title"`
	Author    entities.Author   `json:"author"`
	Category  entities.Category `json:"category"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

type UpdateBookDto struct {
	Title string `json:"title"`
}
