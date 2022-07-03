package main

import (
	"net/http"

	"github.com/matheus-dr/go-test/factories"
)

func main() {
	router := http.NewServeMux()

	authorController := factories.AuthorFactory()

	router.HandleFunc("/author", authorController.Handle)
	router.HandleFunc("/author/id", authorController.HandleParams)

	http.ListenAndServe(":8080", router)
}
