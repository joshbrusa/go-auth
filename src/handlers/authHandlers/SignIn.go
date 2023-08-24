package authHandlers

import (
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
