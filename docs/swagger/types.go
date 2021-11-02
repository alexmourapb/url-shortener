package swagger

type ErrRespBadRequest struct {
	Type  string `json:"type" example:"srn:error:bad_request" validate:"required"`
	Title string `json:"title" example:"required fields are missing" validate:"required"`
}

type ErrRespInternalServerError struct {
	Type  string `json:"type" example:"srn:error:internal_server_error" validate:"required"`
	Title string `json:"title" example:"internal server error"`
}
