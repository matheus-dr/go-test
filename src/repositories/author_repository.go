package repositories

import (
	"errors"
	"github.com/matheus-dr/go-test/src/database"
	"github.com/matheus-dr/go-test/src/entities"
	"time"
)

type AuthorRepository struct {
	fakeDb database.DatabaseStruct
}

var (
	InMemoryAuthorRepository = AuthorRepository{fakeDb: database.Database}
	Authors                  = InMemoryAuthorRepository.fakeDb.AuthorTable.Authors
)

func (r AuthorRepository) CreateAuthor(author entities.Author) {
	InMemoryAuthorRepository.fakeDb.AuthorId += 1

	author.Id = InMemoryAuthorRepository.fakeDb.AuthorId
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

	return entities.Author{}, errors.New("author not found")
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

	return errors.New("author not found")
}

func (r AuthorRepository) DeleteAuthor(id uint) error {
	for i, elem := range Authors {
		if elem.Id == id {
			Authors = append(Authors[:i], Authors[i+1:]...)

			return nil
		}
	}

	return errors.New("author not found")
}
