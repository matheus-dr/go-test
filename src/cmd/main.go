package main

import (
	"github.com/matheus-dr/go-test/src/factories"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	authorController := factories.AuthorFactory()
	categoryController := factories.CategoryFactory()
	bookController := factories.BookFactory()

	router.HandleFunc("/author", authorController.Handle)
	router.HandleFunc("/author/id", authorController.HandleParams)
	router.HandleFunc("/category", categoryController.Handle)
	router.HandleFunc("/category/id", categoryController.HandleParams)
	router.HandleFunc("/book", bookController.Handle)
	router.HandleFunc("/book/id", bookController.HandleParams)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Panic(err)
	}
}
