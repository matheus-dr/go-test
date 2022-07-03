package database

import "github.com/matheus-dr/go-test/entities"

type AuthorTable struct {
	AuthorId uint
	Authors  []entities.Author
}
