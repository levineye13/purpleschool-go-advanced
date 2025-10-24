package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (wrapper *WrapperWriter) WriteHeader(statusCode int) {
	wrapper.ResponseWriter.WriteHeader(statusCode)
	wrapper.StatusCode = statusCode
}
