package link

import (
	"errors"
	"net/http"
	"purpleschool-go/advanced/middleware"
	"purpleschool-go/advanced/pkg/di"
	"purpleschool-go/advanced/pkg/jwt"
	"purpleschool-go/advanced/pkg/req"
	"purpleschool-go/advanced/pkg/res"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	Repo     *LinkRepository
	Jwt      *jwt.JWT
	StatRepo di.IStatRepository
}

type LinkHandler struct {
	baseUrl  string
	Repo     *LinkRepository
	StatRepo di.IStatRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := &LinkHandler{
		baseUrl:  "/links",
		Repo:     deps.Repo,
		StatRepo: deps.StatRepo,
	}

	router.Handle("GET "+linkHandler.baseUrl, middleware.Auth(linkHandler.GetAll(), deps.Jwt))
	router.Handle("GET "+linkHandler.baseUrl+"/count", middleware.Auth(linkHandler.Count(), deps.Jwt))
	router.HandleFunc("GET "+linkHandler.baseUrl+"/{hash}", linkHandler.GoTo())
	router.HandleFunc("POST "+linkHandler.baseUrl, linkHandler.Create())
	router.Handle("PATCH "+linkHandler.baseUrl+"/{id}", middleware.Auth(linkHandler.Update(), deps.Jwt))
	router.HandleFunc("DELETE "+linkHandler.baseUrl+"/{id}", linkHandler.Delete())
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		limit, limitErr := strconv.Atoi(query.Get("limit"))
		offset, offsetErr := strconv.Atoi(query.Get("offset"))

		if limitErr != nil || offsetErr != nil {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		links := handler.Repo.GetAll(limit, offset)

		res.Json(rw, 200, links)
	}
}

func (handler *LinkHandler) Count() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		count := handler.Repo.Count()
		res.Json(rw, 200, count)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")

		if len(hash) == 0 {
			http.Error(rw, "Invalid hash", http.StatusBadRequest)
			return
		}

		link, err := handler.Repo.GetByHash(hash)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}

		handler.StatRepo.AddClick(link.ID)

		http.Redirect(rw, req, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](rw, request)

		if err != nil {
			return
		}

		link := NewLink(body.Url)

		for {
			existedLink, _ := handler.Repo.GetByHash(link.Hash)

			if existedLink == nil {
				break
			}

			link.GenerateHash()
		}

		createdLink, err := handler.Repo.Create(link)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(rw, 201, createdLink)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](rw, request)

		if err != nil {
			return
		}

		id, err := req.ParamToUint(request, "id")

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		updatedLink, err := handler.Repo.Update(&Link{
			Model: gorm.Model{
				ID: *id,
			},
			Url:  body.Url,
			Hash: body.Hash,
		})

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(rw, 200, updatedLink)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		id, err := req.ParamToUint(request, "id")

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		err = handler.Repo.Delete(*id)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(rw, err.Error(), http.StatusNotFound)
				return
			}

			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(rw, 200, nil)
	}
}
