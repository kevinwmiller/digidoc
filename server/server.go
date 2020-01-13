package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/kevinwmiller/digidoc/config"
	log "github.com/kevinwmiller/digidoc/logging"
	"github.com/kevinwmiller/digidoc/server/routes"
	"github.com/kevinwmiller/digidoc/server/routes/auth"
	"github.com/kevinwmiller/digidoc/server/routes/storage"
)

// Server is object that manages routing and middleware
type Server struct {
	cfg    *config.Configuration
	logger *logrus.Logger
	router *mux.Router
	server *http.Server
}

func (s *Server) bindRouter(ctx context.Context, r routes.Router) {
	for route, handler := range r.List() {
		s.router.HandleFunc(route, handler)
	}
}

// New creates a new instance of Server
func New() *Server {
	s := &Server{
		cfg:    config.LoadConfig(".", "config"),
		logger: logrus.New(),
	}
	s.logger.Out = os.Stdout

	if s.cfg.DebugMode {
		s.logger.SetLevel(logrus.DebugLevel)
	} else {
		s.logger.SetLevel(logrus.InfoLevel)
	}

	s.router = mux.NewRouter()

	s.server = &http.Server{
		Handler: s.router,
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Address, s.cfg.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: s.cfg.WriteTimeoutInSeconds,
		ReadTimeout:  s.cfg.ReadTimeoutInSeconds,
	}

	return s
}

// Start launches the Digidoc server. Eventually, this should launch the server in a daemon
func (s *Server) Start() error {

	ctx := context.Background()
	ctx = context.WithValue(ctx, config.Ctx{}, s.cfg)
	ctx = context.WithValue(ctx, log.Ctx{}, s.logger)

	s.router.Use(AddContext(ctx))
	s.router.Use(Logging(ctx))

	s.bindRouter(ctx, &auth.Router{})
	s.bindRouter(ctx, &storage.Router{})
	s.logger.Infof("Starting server on %s\n", s.server.Addr)
	s.logger.Fatal(s.server.ListenAndServe())
	return nil
}
