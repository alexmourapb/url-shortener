package models

type UpdateRequest struct {
	URL    *string `json:"url" validate:"url"`
	Active *bool   `json:"active"`
}
