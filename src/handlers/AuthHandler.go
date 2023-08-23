package handlers

import (
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
