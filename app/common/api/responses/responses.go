package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Details struct {
	Reason   string `json:"reason"`
	Value    string `json:"value"`
	Property string `json:"property"`
}

type FullError struct {
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Details []Details `json:"details,omitempty"`
}

func (f FullError) Error() string {
	return fmt.Sprintf("%+s; %+s; %+v", f.Type, f.Title, f.Details)
}

// shared common responses
var (
	ErrInternalServerError = FullError{Type: "srn:error:internal_server_error", Title: "Internal Server Error"}
	ErrBadRequest          = FullError{Type: "srn:error:bad_request", Title: "Bad ProcessRequest"}
	ErrNotFound            = FullError{Type: "srn:error:not_found", Title: "Not Found"}
	ErrServiceUnavailable  = FullError{Type: "srn:error:service_unavailable", Title: "Service Unavailable"}
	ErrMissingHeaderParams = FullError{Type: "srn:error:missing_header_parameter", Title: "Missing Header Parameter"}
)

func Send(w http.ResponseWriter, response interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
