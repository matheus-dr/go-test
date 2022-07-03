package main

import (
	"net/http"

	"github.com/matheus-dr/go-test/factories"
)

func main() {
	router := http.NewServeMux()

	authorController := factories.AuthorFactory()

	router.HandleFunc("/author", authorController.Handle)

	http.ListenAndServe(":8080", router)
}
