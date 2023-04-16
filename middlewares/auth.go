package middlewares

import (
	"net/http"
	"strings"

	"github.com/sirArthurDayne/rest-ws/helpers"
	"github.com/sirArthurDayne/rest-ws/server"
)

var NO_AUTH_NEEDED = []string{
	"signup", "login",
}

func shouldCheckToken(route string) bool {
	for _, r := range NO_AUTH_NEEDED {
		if strings.Contains(route, r) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(handler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check path before validating token
			if !shouldCheckToken(r.URL.Path) {
				nextHandler.ServeHTTP(w, r)
				return
			}
			// Check JWT token
			_, err := helpers.ValidateJwtAuthToken(s, w, r)
			if err != nil {
				return
			}
			// if passes all checks go to next Handler
			nextHandler.ServeHTTP(w, r)
		})
	}
}
