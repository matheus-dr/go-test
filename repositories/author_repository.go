package repositories

import (
	"time"

	"github.com/matheus-dr/go-test/database"
	"github.com/matheus-dr/go-test/entities"
)

type AuthorRepository struct {
	fakeDb database.DatabaseStruct
}

var (
	IN_MEMORY_AUTHOR_REPOSITORY = AuthorRepository{fakeDb: database.Database}
)

func (r AuthorRepository) CreateAuthor(author entities.Author) {

	IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.AuthorId += 1

	author.Id = IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.AuthorId
	author.CreatedAt = time.Now().UTC()
	author.UpdatedAt = time.Now().UTC()

	IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.AuthorTable.Authors = append(IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.AuthorTable.Authors, author)
}

func (r AuthorRepository) ListAllAuthors() []entities.Author {
	return IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.Authors
}
