package routes

import (
	"net/http"
)

// Index is the main
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}
