package main

import (
	"net/http"
)

func main() {
	router := http.NewServeMux()

	http.ListenAndServe(":8080", router)
}
