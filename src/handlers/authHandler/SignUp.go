package authHandler

import (
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
