package middlewares

import (
	"net/http"
)

// SetMiddlewareJSON set json content type http header
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		resWriter.Header().Set("Content-Type", "application/json")
		next(resWriter, req)
	}
}
