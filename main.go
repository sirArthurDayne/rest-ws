package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirArthurDayne/rest-ws/handlers"
	"github.com/sirArthurDayne/rest-ws/middlewares"
	"github.com/sirArthurDayne/rest-ws/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("[ERROR] cannot load .env file VARS")
	}
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.ServerConfig{
		Port:        PORT,
		JwtSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	// send middlewares
	api.Use(middlewares.CheckAuthMiddleware(s))
	// routes
	router.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	router.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	router.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	api.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts", handlers.InserPostHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)

	router.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
	router.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)
	router.HandleFunc("/ws", s.Hub().HandleWebSockets)
}
