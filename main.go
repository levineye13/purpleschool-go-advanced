package main

import (
	"fmt"
	"net/http"
	"purpleschool-go/advanced/configs"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func main() {
	config, err := configs.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	router := http.NewServeMux()

	router.HandleFunc("/hello", hello)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
