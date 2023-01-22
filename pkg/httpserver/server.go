package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/thehaung/golang-tikv-crud-example/config"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

type HttpServer struct {
	server          *http.Server
	cfg             *config.Config
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, cfg *config.Config) *HttpServer {

	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         cfg.HttpServer.Port,
	}

	server := &HttpServer{
		server:          httpServer,
		cfg:             cfg,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	server.start()
	return server
}

func (s *HttpServer) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *HttpServer) Notify() <-chan error {
	return s.notify
}

func (s *HttpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
