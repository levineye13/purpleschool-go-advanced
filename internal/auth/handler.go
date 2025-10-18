package auth

import (
	"net/http"
	"purpleschool-go/advanced/configs"
	"purpleschool-go/advanced/pkg/req"
	"purpleschool-go/advanced/pkg/res"
)

type AuthHandler struct {
	baseUrl string
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		baseUrl: "/auth",
		Config:  deps.Config,
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

		res.Json(rw, 201, body)
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
