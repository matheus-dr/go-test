package services

import (
	"github.com/matheus-dr/go-test/dto"
	"github.com/matheus-dr/go-test/entities"
	"github.com/matheus-dr/go-test/repositories"
)

type AuthorService struct {
	Repository repositories.AuthorRepository
}

func (s AuthorService) CreateAuthor(input dto.CreateAuthorDto) {
	author := entities.Author{FirstName: input.FirstName, LastName: input.LastName}

	s.Repository.CreateAuthor(author)
}

func (s AuthorService) ListAllAuthors() []dto.ListAuthorDto {
	authors := s.Repository.ListAllAuthors()

	authorsToReturn := make([]dto.ListAuthorDto, len(authors))

	for i, author := range authors {
		authorsToReturn[i] = dto.ListAuthorDto{
			Id:        author.Id,
			FirstName: author.FirstName,
			LastName:  author.LastName,
		}
	}

	return authorsToReturn
}
