package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

func NewJwt(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func checkSecret(secret []byte) error {
	if len(secret) < 32 {
		return &JwtError{
			message: ErrInvalidSecret,
		}
	}

	return nil
}

func (token *JWT) Create(email string) (string, error) {
	secret := []byte(token.Secret)

	err := checkSecret(secret)

	if err != nil {
		return "", err
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"email": email,
	})

	signed, err := newToken.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signed, nil
}
