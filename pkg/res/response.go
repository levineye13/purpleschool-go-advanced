package res

import (
	"encoding/json"
	"net/http"
)

func Json(rw http.ResponseWriter, statusCode int, data any) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	err := json.NewEncoder(rw).Encode(data)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
