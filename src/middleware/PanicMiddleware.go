package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshbrusa/go-auth/src/utils"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error("internal server error")

			if r := recover(); r != nil {
				if panicData, ok := r.(utils.PanicData); ok {
					slog.Error(
						panicData.Msg,
						"file", panicData.File,
						"line", panicData.Line,
					)

					data := map[string]string{
						"msg": panicData.Msg,
					}
					
					json, err := json.Marshal(data)

					if err != nil {
						return
					}
					
					w.Write(json)
				} else {
					slog.Error("Panic data type error.")
				}
			} else {
				slog.Error("Unable to recover panic.")
			}
		}()

		next.ServeHTTP(w, r)
	})
}