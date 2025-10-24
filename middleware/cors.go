package middleware

import (
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		origin := req.Header.Get("Origin")

		if origin == "" {
			next.ServeHTTP(wr, req)
			return
		}

		header := wr.Header()
		header.Set("Access-Control-Allow-Origin", origin)
		header.Set("Access-Control-Allow-Credentials", "true")

		if req.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", "HEAD,GET,POST,PATCH,PUT,DELETE")
			header.Set("Access-Control-Allow-Headers", "accept,content-type,authorization")
			return
		}

		next.ServeHTTP(wr, req)
	})
}
