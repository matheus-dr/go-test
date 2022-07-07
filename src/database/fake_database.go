package database

import (
	"github.com/matheus-dr/go-test/src/entities"
)

type DatabaseStruct struct {
	AuthorTable
	CategoryTable
	BookTable
}

var (
	Database = DatabaseStruct{
		AuthorTable: AuthorTable{
			AuthorId: 0,
			Authors:  make([]entities.Author, 0),
		},
		CategoryTable: CategoryTable{
			CategoryId: 0,
			Categories: make([]entities.Category, 0),
		},
		BookTable: BookTable{
			BookId: 0,
			Books:  make([]entities.Book, 0),
		},
	}
)
