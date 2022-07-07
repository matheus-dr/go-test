package database

import (
	"github.com/matheus-dr/go-test/src/entities"
)

type AuthorTable struct {
	AuthorId uint
	Authors  []entities.Author
}

type CategoryTable struct {
	CategoryId uint
	Categories []entities.Category
}

type BookTable struct {
	BookId uint
	Books  []entities.Book
}
