package req

import (
	"encoding/json"

	"net/http"
)

func Decode[T any](req *http.Request) (*T, error) {
	var body T

	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		return nil, err
	}

	return &body, nil
}
