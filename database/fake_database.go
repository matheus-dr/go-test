package database

import "github.com/matheus-dr/go-test/entities"

type DatabaseStruct struct {
	AuthorTable
}

var (
	Database = DatabaseStruct{
		AuthorTable: AuthorTable{
			AuthorId: 0,
			Authors:  make([]entities.Author, 0),
		},
	}
)
