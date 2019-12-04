package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello home page!")
}

type App struct {
	router *http.ServeMux
	server *http.Server
}

func newApp() *App {
	router := newRouter()
	return &App{
		router: router,
		server: newServer(router),
	}
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", homeHandler)
	return mux
}

func newServer(router *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	app := newApp()
	log.Fatal(app.server.ListenAndServe())
}
