package req

import (
	"net/http"
	"purpleschool-go/advanced/pkg/res"
)

func HandleBody[T any](rw http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := Decode[T](req)

	if err != nil {
		res.Json(rw, 400, err.Error())
		return nil, err
	}

	err = CheckValid(body)

	if err != nil {
		res.Json(rw, 400, err.Error())
		return nil, err
	}

	return body, nil
}
