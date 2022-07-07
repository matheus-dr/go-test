package dto

type CreateBookDto struct {
	Title      string `json:"title"`
	AuthorId   uint   `json:"authorId"`
	CategoryId uint   `json:"categoryId"`
}
