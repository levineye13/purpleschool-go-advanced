package main

import (
	"fmt"
	"net/http"
	"purpleschool-go/advanced/configs"
	"purpleschool-go/advanced/internal/auth"
)

func main() {
	config, err := configs.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: router,
	}

	err = server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
