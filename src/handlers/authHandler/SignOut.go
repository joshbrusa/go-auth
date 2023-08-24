package authHandler

import (
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
