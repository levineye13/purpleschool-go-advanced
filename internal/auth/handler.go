package auth

import (
	"net/http"
	"purpleschool-go/advanced/configs"
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
		body, err := req.HandleBody[TRegisterRequest](rw, request)

		if err != nil {
			res.Json(rw, 400, err.Error())
			return
		}

		userEmail, err := handler.AuthService.Register(body.Email, body.Name, body.Password)

		if err != nil {
			res.Json(rw, 400, err.Error())
			return
		}

		res.Json(rw, 201, userEmail)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[TLoginRequest](rw, request)

		if err != nil {
			res.Json(rw, 400, err.Error())
			return
		}

		res.Json(rw, 200, body)
	}
}
