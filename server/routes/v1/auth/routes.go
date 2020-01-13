package auth

import (
	"net/http"
)

// Router manages routes related to authentication of user objects
type Router struct{}

// List returns a list of authentication route handlers
func (r Router) List() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/auth/login":  login,
		"/auth/logout": logout,
	}
}
