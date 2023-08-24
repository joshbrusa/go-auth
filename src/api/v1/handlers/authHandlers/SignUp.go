package authHandlers

import (
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
