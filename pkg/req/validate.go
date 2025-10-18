package req

import "github.com/go-playground/validator/v10"

func CheckValid[T any](body T) error {
	validator := validator.New()
	err := validator.Struct(body)

	return err
}
