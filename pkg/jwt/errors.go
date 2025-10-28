package jwt

type JwtError struct {
	message string
}

const (
	ErrInvalidSecret = "invalid secret"
	ErrInvalidToken  = "invalid secret"
)

func (jwtError *JwtError) Error() string {
	return jwtError.message
}
