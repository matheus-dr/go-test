package dto

import "time"

type CreateCategoryDto struct {
	Description string `json:"description"`
}

type ListCategoryDto struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
}

type FindCategoryDto struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateCategoryDto struct {
	Description string `json:"description"`
}
