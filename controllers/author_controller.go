package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matheus-dr/go-test/dto"
	"github.com/matheus-dr/go-test/services"
)

type AuthorController struct {
	service services.AuthorService
}

func (c AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var data dto.CreateAuthorDto
	json.Unmarshal(input, &data)

	c.service.CreateAuthor(data)
}
