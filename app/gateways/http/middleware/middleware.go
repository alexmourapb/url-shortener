package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/alexmourapb/url-shortener/app/common/shared"
)

// Cors applies cors rules to router
func Cors(r *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(r)
}

func AddHeaders(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// x-request-id
	requestID := r.Header.Get("x-request-id")
	if requestID == "" {
		requestID = uuid.New().String()
	}
	ctx := context.WithValue(r.Context(), shared.KeyRequestID, requestID)

	r = r.WithContext(ctx)
	w.Header().Set("x-request-id", requestID)
	next.ServeHTTP(w, r)
}

// Removes the trailing slash from request, except if it is the root url.
// If the url is https://teste.com.br/api or https://teste.com.br/api/
// both will match.
// This was done as gorilla mux default method for this doesn't support POST requests: https://github.com/gorilla/mux/issues/30
// Usage:
// n := negroni.Classic()
// n.UseFunc(middleware.TrimSlashSuffix)
func TrimSlashSuffix(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path != "/" {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	}

	next.ServeHTTP(w, r)
}
