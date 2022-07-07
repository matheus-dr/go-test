package services

import (
	"github.com/matheus-dr/go-test/src/dto"
	"github.com/matheus-dr/go-test/src/entities"
	"github.com/matheus-dr/go-test/src/repositories"
	"log"
)

type BookService struct {
	BookRepository     repositories.BookRepository
	AuthorRepository   repositories.AuthorRepository
	CategoryRepository repositories.CategoryRepository
}

func (s BookService) CreateBook(input dto.CreateBookDto) {
	author, err := s.AuthorRepository.FindOneAuthor(input.AuthorId)
	if err != nil {
		log.Panic(err)
	}

	category, err := s.CategoryRepository.FindOneCategory(input.CategoryId)
	if err != nil {
		log.Panic(err)
	}

	book := entities.Book{
		Title:    input.Title,
		Author:   author,
		Category: category,
	}

	s.BookRepository.CreateBook(book)
}
