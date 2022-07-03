package services

import (
	"github.com/matheus-dr/go-test/dto"
	"github.com/matheus-dr/go-test/repositories"
)

type AuthorService struct {
	Repository repositories.AuthorRepository
}

func (a AuthorService) CreateAuthor(author dto.CreateAuthorDto) {
	a.Repository.CreateAuthor(author)
}
