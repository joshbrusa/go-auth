package middleware

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			w.WriteHeader(http.StatusInternalServerError)

			rec := recover()

			if rec == nil || reflect.TypeOf(rec).String() != "string" {
				return
			}

			json, err := json.Marshal(rec)

			if err != nil {
				return
			}

			// fix this response, json not working
			w.Write(json)
		}()

		next.ServeHTTP(w, r)
	})
}