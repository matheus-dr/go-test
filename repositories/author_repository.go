package repositories

import (
	"time"

	"github.com/matheus-dr/go-test/entities"
)

type AuthorRepository struct {
	authorId uint
	Authors  []entities.Author
}

func (r AuthorRepository) CreateAuthor(author entities.Author) {
	r.authorId += 1

	author.Id = r.authorId
	author.CreatedAt = time.Now().UTC()
	author.UpdatedAt = time.Now().UTC()

	r.Authors = append(r.Authors, author)
}

func (r AuthorRepository) ListAllAuthors() []entities.Author {
	return r.Authors
}
