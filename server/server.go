package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kevinwmiller/digidoc/log"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/kevinwmiller/digidoc/config"
	"github.com/kevinwmiller/digidoc/server/middleware"
	"github.com/kevinwmiller/digidoc/server/routes"
)

type Server struct {
	router *mux.Router
	server *http.Server
}

type ServerOptions struct {
	Port int
}

func (s *Server) Start(options *ServerOptions) error {

	ctx := context.Background()
	cfg := config.LoadConfig(".", "config")
	logger := logrus.New()
	logger.Out = os.Stdout

	if cfg.DebugMode {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	ctx = context.WithValue(ctx, config.Ctx{}, &cfg)
	ctx = context.WithValue(ctx, log.Ctx{}, &logger)

	s.router = mux.NewRouter()
	s.router.HandleFunc("/login", routes.Login)
	s.router.HandleFunc("/logout", routes.Logout)
	s.router.HandleFunc("/upload", routes.Upload)
	s.router.HandleFunc("/download", routes.Download)
	s.router.HandleFunc("/delete", routes.Delete)
	s.router.HandleFunc("/", routes.Index)
	// s.router.HandleFunc("/{id}", routes.View)
	http.Handle("/", middleware.AddContext(ctx, s.router))

	s.server = &http.Server{
		Handler: s.router,
		Addr:    fmt.Sprintf("127.0.0.1:%d", options.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(s.server.ListenAndServe())
	return nil
}
