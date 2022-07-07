package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/matheus-dr/go-test/src/dto"
	"github.com/matheus-dr/go-test/src/services"
	"io/ioutil"
	"log"
	"net/http"
)

type BookController struct {
	Service services.BookService
}

func (c BookController) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.createBook(w, r)

	default:
		fmt.Println("Unknown method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c BookController) createBook(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var data dto.CreateBookDto
	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Panic(err)
	}

	c.Service.CreateBook(data)

	w.WriteHeader(http.StatusCreated)
}
