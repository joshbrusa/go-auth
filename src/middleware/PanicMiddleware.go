package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()

			if r == nil {
				return
			}

			w.WriteHeader(http.StatusInternalServerError)

			msg := "Internal server error."

			slog.Error(msg)

			data := map[string]string{
				"msg": msg,
			}

			json, err := json.Marshal(data)

			if err != nil {
				slog.Error(err.Error())
				return
			}

			w.Write(json)


			msg, ok := r.(string)

			if !ok {
				slog.Error("Wrong panic type.")
			}

			slog.Error(msg)
		}()

		next.ServeHTTP(w, r)
	})
}
