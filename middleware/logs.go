package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		writer := &WrapperWriter{
			ResponseWriter: rw,
		}

		next.ServeHTTP(writer, req)

		log.Println(writer.StatusCode, req.Method, req.URL.Path, time.Since(start))
	})
}
