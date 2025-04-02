package dto

type PostCreateReuqest struct {
	Title string `json:"title" validate:"required"`
	Desc  string `json:"desc" validate:"required"`
	Image string `json:"image" validate:"required"`
}

type PostListResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
}

type PostDetailResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
}
