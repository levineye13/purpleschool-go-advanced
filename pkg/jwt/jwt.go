package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}
type JWTData struct {
	Email string
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

func (token *JWT) Create(jwtData *JWTData) (string, error) {
	secret := []byte(token.Secret)

	err := checkSecret(secret)

	if err != nil {
		return "", err
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"email": jwtData.Email,
	})

	signed, err := newToken.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signed, nil
}

func (jwtToken *JWT) Parse(token string) (bool, *JWTData) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(jwtToken.Secret), nil
	})

	if err != nil {
		return false, nil
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	email, ok := claims["email"].(string)

	if !ok {
		return false, nil
	}

	return parsedToken.Valid, &JWTData{
		Email: email,
	}
}
