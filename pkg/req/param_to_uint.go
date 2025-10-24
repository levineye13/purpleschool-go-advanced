package req

import (
	"net/http"
	"strconv"
)

func ParamToUint(request *http.Request, param string) (*uint, error) {
	paramUint64, err := strconv.ParseUint(request.PathValue(param), 10, 32)

	if err != nil {
		return nil, err
	}

	paramUint := uint(paramUint64)

	return &paramUint, nil
}
