package middleware

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

// AddContext adds the given context to any nested http handlers
func AddContext(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler")
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
