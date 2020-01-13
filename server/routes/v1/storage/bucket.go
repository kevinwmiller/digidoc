package storage

import "net/http"

func delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
