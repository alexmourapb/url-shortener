package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"

	"github.com/alexmourapb/url-shortener/app/common/shared"
)

var rootUrl = "/"
var apiUrl = "/shortener-api"

func Test_TrimSlashSuffix(t *testing.T) {
	tests := []struct {
		url  string
		name string
	}{
		{
			url:  rootUrl,
			name: "Test root url",
		},
		{
			url:  apiUrl,
			name: "Test url without trailing slash",
		},
		{
			url:  apiUrl + "/",
			name: "Test url with trailing slash",
		},
	}

	n := negroni.New()
	n.UseFunc(TrimSlashSuffix)

	r := mux.NewRouter()

	responseOk := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	r.HandleFunc(apiUrl, responseOk).Methods(http.MethodGet)
	r.HandleFunc(rootUrl, responseOk).Methods(http.MethodGet)

	n.UseHandler(r)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			n.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
}

func Test_AddHeaders(t *testing.T) {
	n := negroni.New()
	n.UseFunc(AddHeaders)

	r := mux.NewRouter()

	responseOk := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "abc123", w.Header().Get("x-request-id"))

		assert.Equal(t, "abc123", r.Context().Value(shared.KeyRequestID))
		w.WriteHeader(http.StatusOK)
	}

	r.HandleFunc(rootUrl, responseOk).Methods(http.MethodGet)

	n.UseHandler(r)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	req.Header.Set("x-request-id", "abc123")

	w := httptest.NewRecorder()
	n.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
