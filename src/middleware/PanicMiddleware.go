package middleware

import (
	"log/slog"
	"net/http"

	"github.com/joshbrusa/go-auth/src/utils"
)

/*
Reminder: Panic middleware runs at the end of every request.
*/
func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rec *http.Request) {
		defer func() {
			rec := recover()

			if rec == nil {
				return
			}

			msg := "Internal server error."

			slog.Error(msg)

			response := &utils.Response{
				ResponseWriter: w,
				StatusCode: http.StatusInternalServerError,
				Data: msg,
			}

			response.Send()

			recMsg, ok := rec.(string)

			if !ok {
				slog.Error("Wrong panic type.")
				return
			}

			slog.Error(recMsg)
		}()

		next.ServeHTTP(w, rec)
	})
}
