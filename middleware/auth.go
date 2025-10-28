package middleware

import (
	"context"
	"net/http"
	"purpleschool-go/advanced/pkg/jwt"
	"strings"
)

type Key string

const (
	AuthContextKey Key = "AuthContextKey"
)

func writeUnauthorized(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusUnauthorized)
	rw.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func Auth(next http.Handler, jwt *jwt.JWT) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		auth := req.Header.Get("Authorization")

		if auth == "" {
			writeUnauthorized(rw)
			return
		}

		splitedAuth := strings.Split(auth, " ")

		if splitedAuth[0] != "Bearer" || splitedAuth[1] == "" {
			writeUnauthorized(rw)
			return
		}

		isValid, jwtData := jwt.Parse(splitedAuth[1])

		if !isValid {
			writeUnauthorized(rw)
			return
		}

		authContext := context.WithValue(req.Context(), AuthContextKey, jwtData.Email)
		reqWithAuthContext := req.WithContext(authContext)

		next.ServeHTTP(rw, reqWithAuthContext)
	})
}
