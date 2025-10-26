package auth

import (
	"errors"
	"net/http"
	"purpleschool-go/advanced/configs"
	"purpleschool-go/advanced/pkg/jwt"
	"purpleschool-go/advanced/pkg/req"
	"purpleschool-go/advanced/pkg/res"
)

type AuthHandler struct {
	*configs.Config
	baseUrl     string
	AuthService *AuthService
}
type AuthHandlerDeps struct {
	*configs.Config
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		baseUrl:     "/auth",
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST "+handler.baseUrl+"/register", handler.Register())
	router.HandleFunc("POST "+handler.baseUrl+"/login", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[RegisterRequest](rw, request)

		if err != nil {
			res.Json(rw, http.StatusUnauthorized, err.Error())
			return
		}

		newJwt, err := handler.AuthService.Register(body.Email, body.Name, body.Password)

		if err != nil {
			res.Json(rw, http.StatusUnauthorized, err.Error())
			return
		}

		registerRes := &RegisterResponse{
			Token: newJwt,
		}

		res.Json(rw, 201, registerRes)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LoginRequest](rw, request)

		if err != nil {
			res.Json(rw, http.StatusUnauthorized, err.Error())
			return
		}

		token, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			var jwtErr *jwt.JwtError

			if errors.As(err, &jwtErr) {
				res.Json(rw, http.StatusInternalServerError, "internal server error")
				return
			}

			res.Json(rw, http.StatusUnauthorized, err.Error())
			return
		}

		loginRes := &LoginResponse{
			Token: token,
		}

		res.Json(rw, 200, loginRes)
	}
}
