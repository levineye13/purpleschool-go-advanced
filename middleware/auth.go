package middleware

import (
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		auth := req.Header.Get("Authorization")

		if auth == "" {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		splitedAuth := strings.Split(auth, " ")

		if splitedAuth[0] != "Bearer" || splitedAuth[1] == "" {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, req)
	})
}
