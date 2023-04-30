package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirArthurDayne/rest-ws/database"
	"github.com/sirArthurDayne/rest-ws/repository"
	"github.com/sirArthurDayne/rest-ws/websocket"
)

type ServerConfig struct {
	Port        string
	JwtSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *ServerConfig
	Hub() *websocket.Hub
}

// broker: handles all servers
type Broker struct {
	config *ServerConfig
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *ServerConfig {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *ServerConfig) (*Broker, error) {
	// error checking for serverconfig
	if config.Port == "" {
		return nil, errors.New("[ERROR]: Port is required")
	}
	if config.JwtSecret == "" {
		return nil, errors.New("[ERROR]: JWT Secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("[ERROR]: DatabaseUrl is required")
	}
	// make broker
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, router *mux.Router)) {
	// create router and bind Handlers
	b.router = mux.NewRouter()
	binder(b, b.router)
	// allow connection from outside localhost
	handler := cors.Default().Handler(b.router)
	// set DB repository
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	// websocket run
	go b.hub.Run()

	// initialize models
	repository.SetRepository(repo)
	// initialize server
	log.Println("Starting server on port ", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("[ERROR] cannot starte server: ", err)
	} else {
        log.Fatal("[ERROR] server stopped!")
    }
}
