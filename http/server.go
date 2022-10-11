package http

import (
	"log"
	"net/http"
)

func StartServer() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: Init(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
