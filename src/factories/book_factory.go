package factories

import (
	"github.com/matheus-dr/go-test/src/controllers"
	"github.com/matheus-dr/go-test/src/repositories"
	"github.com/matheus-dr/go-test/src/services"
)

func BookFactory() controllers.BookController {
	bookRepository := repositories.BookRepository{}
	authorRepository := repositories.AuthorRepository{}
	categoryRepository := repositories.CategoryRepository{}

	service := services.BookService{
		BookRepository:     bookRepository,
		AuthorRepository:   authorRepository,
		CategoryRepository: categoryRepository,
	}

	return controllers.BookController{Service: service}
}
