package main

import (
	"fmt"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/hello", hello)

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
