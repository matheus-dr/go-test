package dto

import "time"

type CreateAuthorDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type ListAuthorDto struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type FindAuthorDto struct {
	Id        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateAuthorDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
