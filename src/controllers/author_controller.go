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

type AuthorController struct {
	Service services.AuthorService
}

func (c AuthorController) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.createAuthor(w, r)

	case http.MethodGet:
		c.listAllAuthors(w, r)

	default:
		fmt.Println("Unknown method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c AuthorController) HandleParams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.findOneAuthor(w, r)

	case http.MethodPut:
		c.updateAuthor(w, r)

	case http.MethodDelete:
		c.deleteAuthor(w, r)

	default:
		fmt.Println("Unknown method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c AuthorController) createAuthor(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var data dto.CreateAuthorDto
	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Panic(err)
	}

	c.Service.CreateAuthor(data)

	w.WriteHeader(http.StatusCreated)
}

func (c AuthorController) listAllAuthors(w http.ResponseWriter, r *http.Request) {
	output := c.Service.ListAllAuthors()

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Panic(err)
	}
}

func (c AuthorController) findOneAuthor(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	output := c.Service.FindOneAuthor(uint(id))

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Panic(err)
	}
}

func (c AuthorController) updateAuthor(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	var data dto.UpdateAuthorDto
	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Panic(err)
	}

	c.Service.UpdateAuthor(uint(id), data)

	w.WriteHeader(http.StatusOK)
}

func (c AuthorController) deleteAuthor(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	c.Service.DeleteAuthor(uint(id))

	w.WriteHeader(http.StatusOK)
}
