package main

import (
	"fmt"
	"net/http"
	"purpleschool-go/advanced/configs"
	"purpleschool-go/advanced/internal/auth"
	"purpleschool-go/advanced/internal/link"
	"purpleschool-go/advanced/middleware"
	"purpleschool-go/advanced/pkg/db"
)

func main() {
	config, err := configs.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	db := db.NewDb(config)

	linkRepository := link.NewLinkRepository(db)

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		Repo: linkRepository,
	})

	middlewares := middleware.Chain(
		middleware.Cors,
		middleware.Logger,
	)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: middlewares(router),
	}

	err = server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
