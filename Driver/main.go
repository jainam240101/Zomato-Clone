package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/jainam240101/zomato-clone/Driver/Handlers"
)

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tracking", handlers.Tracking)
	mux.HandleFunc("/search", handlers.Search)
	return mux
}

func main() {
	handlers.NewHandlers()
	server := http.Server{
		Addr:    fmt.Sprint(":8000"),
		Handler: NewHandler(),
	}
	// Run server
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed ! ")
	}

}
