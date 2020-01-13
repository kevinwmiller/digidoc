package storage

import "net/http"

func download(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download"))
}
