package routes

import "net/http"

func Upload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload"))
}
