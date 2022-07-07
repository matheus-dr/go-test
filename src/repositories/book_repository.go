package repositories

import (
	"errors"
	"github.com/matheus-dr/go-test/src/database"
	"github.com/matheus-dr/go-test/src/entities"
	"time"
)

type BookRepository struct {
	fakeDb database.DatabaseStruct
}

var (
	InMemoryBookRepository = BookRepository{fakeDb: database.Database}
	Books                  = InMemoryBookRepository.fakeDb.BookTable.Books
)

func (r BookRepository) CreateBook(book entities.Book) {
	InMemoryBookRepository.fakeDb.BookId += 1

	book.Id = InMemoryBookRepository.fakeDb.BookId
	book.CreatedAt = time.Now().UTC()
	book.UpdatedAt = time.Now().UTC()

	Books = append(Books, book)
}

func (r BookRepository) FindAllBooks() []entities.Book {
	return Books
}

func (r BookRepository) FindOneBook(id uint) (entities.Book, error) {
	for _, book := range Books {
		if book.Id == id {
			return book, nil
		}
	}

	return entities.Book{}, errors.New("book not found")
}

func (r BookRepository) UpdateBook(id uint, data entities.Book) error {
	for i, book := range Books {
		if book.Id == id {
			if data.Title != "" {
				book.Title = data.Title
			}

			Books[i] = book

			return nil
		}
	}

	return errors.New("book not found")
}

func (r BookRepository) DeleteBook(id uint) error {
	for i, book := range Books {
		if book.Id == id {
			Books = append(Books[:i], Books[i+1:]...)

			return nil
		}
	}

	return errors.New("book not found")
}
