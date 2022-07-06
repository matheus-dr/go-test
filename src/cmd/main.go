package main

import (
	"github.com/matheus-dr/go-test/src/factories"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	authorController := factories.AuthorFactory()
	categoryController := factories.CategoryFactory()

	router.HandleFunc("/author", authorController.Handle)
	router.HandleFunc("/author/id", authorController.HandleParams)
	router.HandleFunc("/category", categoryController.Handle)
	router.HandleFunc("/category/id", categoryController.HandleParams)

	http.ListenAndServe(":8080", router)
}
