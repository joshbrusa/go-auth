package userhandlers

import "net/http"

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}