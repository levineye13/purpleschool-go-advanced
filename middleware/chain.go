package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for index := len(middlewares) - 1; index >= 0; index-- {
			next = middlewares[index](next)
		}

		return next
	}
}
