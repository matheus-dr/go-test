package repositories

import (
	"github.com/matheus-dr/go-test/dto"
	"github.com/matheus-dr/go-test/entities"
)

type AuthorRepository struct {
	authorId int
	Authors  []entities.Author
}

func (a AuthorRepository) CreateAuthor(input dto.CreateAuthorDto) {
	a.authorId += 1
	print(a.authorId, "author created")
	// TODO: check how create an author in slice
}
