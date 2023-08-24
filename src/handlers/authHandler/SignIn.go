package authHandler

import (
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
