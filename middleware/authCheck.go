package middleware

import (
	"net/http"
)

func AuthCheck(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

        handler(w, r)
	}
}