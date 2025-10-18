package req

import "net/http"

func HandleBody[T any](rw http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := Decode[T](req)

	if err != nil {
		return nil, err
	}

	err = CheckValid(body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
