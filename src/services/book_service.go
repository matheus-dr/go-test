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

func (s BookService) ListAllBooks() []dto.ListBookDto {
	books := s.BookRepository.FindAllBooks()

	booksToReturn := make([]dto.ListBookDto, len(books))

	for i, book := range books {
		booksToReturn[i] = dto.ListBookDto{
			Id:       book.Id,
			Title:    book.Title,
			Author:   book.Author,
			Category: book.Category,
		}
	}

	return booksToReturn
}

func (s BookService) FindOneBook(id uint) dto.FindBookDto {
	book, err := s.BookRepository.FindOneBook(id)
	if err != nil {
		log.Panic(err)
	}

	return dto.FindBookDto(book)
}

func (s BookService) UpdateBook(id uint, data dto.UpdateBookDto) {
	input := entities.Book{Title: data.Title}

	err := s.BookRepository.UpdateBook(id, input)
	if err != nil {
		log.Panic(err)
	}
}

func (s BookService) DeleteBook(id uint) {
	err := s.BookRepository.DeleteBook(id)
	if err != nil {
		log.Panic(err)
	}
}
