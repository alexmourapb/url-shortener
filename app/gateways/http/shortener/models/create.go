package models

type CreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type CreateResponse struct {
	URL string `json:"short_url"`
}
