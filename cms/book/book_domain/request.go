package book_domain

type BookRequestCreate struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Cover  string `json:"cover"`
}
