package auth

import (
	"errors"
	"purpleschool-go/advanced/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, name, password string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	newUser := &user.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	_, err = service.UserRepository.Create(newUser)

	if err != nil {
		return "", err
	}

	return newUser.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	user, _ := service.UserRepository.FindByEmail(email)

	if user == nil {
		return "", errors.New(ErrUserNotFound)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New(ErrUserNotFound)
	}

	return user.Email, nil
}
