package routes

import "net/http"

func Download(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download"))
}
