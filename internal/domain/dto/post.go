package dto

type PostCreateReuqest struct {
	Title string `json:"title" validate:"required"`
	Desc  string `json:"desc" validate:"required"`
	Image string `json:"image" validate:"required"`
}
