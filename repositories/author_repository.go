package repositories

import (
	"github.com/matheus-dr/go-test/dto"
	"github.com/matheus-dr/go-test/entities"
)

type AuthorRepository struct {
	authorId int
	authors  []entities.Author
}

func (a AuthorRepository) CreateAuthor(input dto.CreateAuthorDto) {
	a.authorId += 1
	print(a.authorId, "author created")
}
