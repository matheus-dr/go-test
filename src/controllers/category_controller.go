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

type CategoryController struct {
	Service services.CategoryService
}

func (c CategoryController) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.createCategory(w, r)

	case http.MethodGet:
		c.listAllCategories(w, r)

	default:
		fmt.Println("Unknown method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c CategoryController) HandleParams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.findOneCategory(w, r)

	case http.MethodPut:
		c.updateCategory(w, r)

	case http.MethodDelete:
		c.deleteCategory(w, r)

	default:
		fmt.Println("Unknown method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c CategoryController) createCategory(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var data dto.CreateCategoryDto
	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Panic(err)
	}

	c.Service.CreateCategory(data)

	w.WriteHeader(http.StatusCreated)
}

func (c CategoryController) listAllCategories(w http.ResponseWriter, r *http.Request) {
	output := c.Service.ListAllCategories()

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Panic(err)
	}
}

func (c CategoryController) findOneCategory(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	output := c.Service.FindOneCategory(uint(id))

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Panic(err)
	}
}

func (c CategoryController) updateCategory(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	var data dto.UpdateCategoryDto
	err = json.Unmarshal(input, &data)
	if err != nil {
		log.Panic(err)
	}

	c.Service.UpdateCategory(uint(id), data)

	w.WriteHeader(http.StatusOK)
}

func (c CategoryController) deleteCategory(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Panic(err)
	}

	c.Service.DeleteCategory(uint(id))

	w.WriteHeader(http.StatusOK)
}
