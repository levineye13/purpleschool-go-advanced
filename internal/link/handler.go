package link

import (
	"net/http"
	"purpleschool-go/advanced/pkg/req"
	"purpleschool-go/advanced/pkg/res"
)

type LinkHandlerDeps struct {
	Repo *LinkRepository
}

type LinkHandler struct {
	baseUrl string
	Repo    *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := &LinkHandler{
		baseUrl: "/links",
		Repo:    deps.Repo,
	}

	router.HandleFunc("GET "+linkHandler.baseUrl, linkHandler.GetAll())
	router.HandleFunc("GET "+"/{alias}", linkHandler.GoTo())
	router.HandleFunc("POST "+linkHandler.baseUrl, linkHandler.Create())
	router.HandleFunc("PATCH "+linkHandler.baseUrl, linkHandler.Update())
	router.HandleFunc("DELETE "+linkHandler.baseUrl+"/{id}", linkHandler.Delete())
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		res.Json(rw, 200, []any{})
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		res.Json(rw, 200, []any{})
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](rw, request)

		if err != nil {
			res.Json(rw, 400, err.Error())
		}

		res.Json(rw, 201, body)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](rw, request)

		if err != nil {
			res.Json(rw, 400, err.Error())
		}

		res.Json(rw, 200, body)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		id := request.PathValue("id")

		if id == "" {
		}

	}
}
