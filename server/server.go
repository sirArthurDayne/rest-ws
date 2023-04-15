package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Port        string
	JwtSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *ServerConfig
}

// broker: handles all servers
type Broker struct {
	config *ServerConfig
	router mux.Router
}

func (b *Broker) Config() *ServerConfig {
	return b.config
}

func NewServer(ctx context.Context, config *ServerConfig) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("[ERROR]: Port is required")
	}
	if config.JwtSecret == "" {
		return nil, errors.New("[ERROR]: JWT Secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("[ERROR]: DatabaseUrl is required")
	}
	broker := &Broker{
		config: config,
		router: *mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, router mux.Router)) {
	b.router = *mux.NewRouter()
	binder(b, b.router)
	log.Println("Starting server on port ", b.Config().Port)
	if err := http.ListenAndServe(b.Config().Port, &b.router); err != nil {
        log.Fatal("ListenAndServe: ", err)
	}
}
