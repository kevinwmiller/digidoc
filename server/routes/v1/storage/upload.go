package storage

import "net/http"

func upload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload"))
}
