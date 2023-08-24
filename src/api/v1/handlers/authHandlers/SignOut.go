package authHandlers

import (
	"net/http"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
