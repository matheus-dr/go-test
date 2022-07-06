package services

import (
	"github.com/matheus-dr/go-test/src/dto"
	"github.com/matheus-dr/go-test/src/entities"
	"github.com/matheus-dr/go-test/src/repositories"
	"log"
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

func (s AuthorService) FindOneAuthor(id uint) dto.FindAuthorDto {
	author, err := s.Repository.FindOneAuthor(id)
	if err != nil {
		log.Panic(err)
	}

	return dto.FindAuthorDto(author)
}

func (s AuthorService) UpdateAuthor(id uint, data dto.UpdateAuthorDto) {
	input := entities.Author{FirstName: data.FirstName, LastName: data.LastName}

	err := s.Repository.UpdateAuthor(id, input)
	if err != nil {
		log.Panic(err)
	}
}

func (s AuthorService) DeleteAuthor(id uint) {
	err := s.Repository.DeleteAuthor(id)
	if err != nil {
		log.Panic(err)
	}
}
