package repositories

import (
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
