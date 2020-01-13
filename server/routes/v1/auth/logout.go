package auth

import "net/http"

func logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
