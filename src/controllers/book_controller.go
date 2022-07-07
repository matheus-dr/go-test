package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/matheus-dr/go-test/src/dto"
	"github.com/matheus-dr/go-test/src/services"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type BookController struct {
	Service services.BookService
}

func (c BookController) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.deleteBook(w, r)

	case http.MethodGet:
		c.listAllBooks(w, r)

	default:
		fmt.Println("Unknown method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c BookController) HandleParams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.deleteBook(w, r)

	case http.MethodPut:
		c.deleteBook(w, r)

	case http.MethodDelete:
		c.deleteBook(w, r)

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

func (c BookController) listAllBooks(w http.ResponseWriter, r *http.Request) {
	output := c.Service.ListAllBooks()

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Panic(err)
	}
}

func (c BookController) findOneBook(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	output := c.Service.FindOneBook(uint(id))

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Panic(err)
	}
}

func (c BookController) updateBook(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	var data dto.UpdateBookDto
	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Panic(err)
	}

	c.Service.UpdateBook(uint(id), data)

	w.WriteHeader(http.StatusOK)
}

func (c BookController) deleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	c.Service.DeleteBook(uint(id))

	w.WriteHeader(http.StatusOK)
}
