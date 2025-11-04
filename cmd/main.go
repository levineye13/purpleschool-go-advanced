package main

import (
	"fmt"
	"net/http"
	"purpleschool-go/advanced/configs"
	"purpleschool-go/advanced/internal/auth"
	"purpleschool-go/advanced/internal/link"
	"purpleschool-go/advanced/internal/stat"
	"purpleschool-go/advanced/internal/user"
	"purpleschool-go/advanced/middleware"
	"purpleschool-go/advanced/pkg/db"
	"purpleschool-go/advanced/pkg/event"
	"purpleschool-go/advanced/pkg/jwt"
)

func main() {
	config, err := configs.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	db := db.NewDb(config)
	jwt := jwt.NewJwt(config.Auth.Secret)
	eventBus := event.NewEventBus()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	authService := auth.NewAuthService(userRepository, jwt)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		StatRepo: statRepository,
		EventBus: eventBus,
	})

	go statService.AddClick()

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		Repo:     linkRepository,
		Jwt:      jwt,
		EventBus: eventBus,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatService: statService,
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
