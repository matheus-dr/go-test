package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matheus-dr/go-test/dto"
	"github.com/matheus-dr/go-test/services"
)

type AuthorController struct {
	Service services.AuthorService
}

func (c AuthorController) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateAuthor(w, r)

		// TODO: change this to GET AND Receive a param that is an ID
	case http.MethodGet:
		c.ListAllAuthors(w, r)

	default:
		fmt.Println("Unknown method:", r.Method)
	}
}

func (c AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var data dto.CreateAuthorDto
	json.Unmarshal(input, &data)

	c.Service.CreateAuthor(data)

	w.WriteHeader(http.StatusCreated)
}

func (c AuthorController) ListAllAuthors(w http.ResponseWriter, r *http.Request) {
	output := c.Service.ListAllAuthors()

	json.NewEncoder(w).Encode(output)
}
