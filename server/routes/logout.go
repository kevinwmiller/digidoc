package routes

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
