package repositories

import (
	"errors"
	"time"

	"github.com/matheus-dr/go-test/database"
	"github.com/matheus-dr/go-test/entities"
)

type AuthorRepository struct {
	fakeDb database.DatabaseStruct
}

var (
	IN_MEMORY_AUTHOR_REPOSITORY = AuthorRepository{fakeDb: database.Database}
	Authors                     = IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.AuthorTable.Authors
)

func (r AuthorRepository) CreateAuthor(author entities.Author) {
	IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.AuthorId += 1

	author.Id = IN_MEMORY_AUTHOR_REPOSITORY.fakeDb.AuthorId
	author.CreatedAt = time.Now().UTC()
	author.UpdatedAt = time.Now().UTC()

	Authors = append(Authors, author)
}

func (r AuthorRepository) ListAllAuthors() []entities.Author {
	return Authors
}

func (r AuthorRepository) FindOneAuthor(id uint) (entities.Author, error) {
	for _, author := range Authors {
		if author.Id == id {
			return author, nil
		}
	}

	return entities.Author{}, errors.New("Author not found")
}

func (r AuthorRepository) UpdateAuthor(id uint, input entities.Author) error {
	for i, author := range Authors {
		if author.Id == id {
			if input.FirstName != "" {
				author.FirstName = input.FirstName
			}

			if input.LastName != "" {
				author.LastName = input.LastName
			}

			author.UpdatedAt = time.Now().UTC()

			Authors[i] = author

			return nil
		}
	}

	return errors.New("Author not found")
}

func (r AuthorRepository) DeleteAuthor(id uint) error {
	for i, elem := range Authors {
		if elem.Id == id {
			Authors = append(Authors[:i], Authors[i+1:]...)

			return nil
		}
	}

	return errors.New("Author not found")
}
