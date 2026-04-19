package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/durianpay/fullstack-boilerplate/internal/middleware"
	"github.com/durianpay/fullstack-boilerplate/internal/openapigen"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	oapinethttpmw "github.com/oapi-codegen/nethttp-middleware"
)

type Server struct {
	router  http.Handler
	service *http.Server
}

const (
	readTimeout  = 10
	writeTimeout = 10
	idleTimeout  = 60
)

func NewServer(apiHandler openapigen.ServerInterface, openapiYamlPath string, appEnv string, jwtSecret []byte) *Server {
	swagger, err := openapigen.GetSwagger()
	if err != nil {
		log.Fatalf("failed to load swagger: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.CORS)

	if appEnv == "development" {
		r.Use(middleware.Logger)
	}

	r.Route("/", func(api chi.Router) {
		api.Use(middleware.Auth(swagger, jwtSecret))
		api.Use(oapinethttpmw.OapiRequestValidatorWithOptions(
			swagger,
			&oapinethttpmw.Options{
				DoNotValidateServers:  true,
				SilenceServersWarning: true,
				Options: openapi3filter.Options{
					AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
				},
			},
		))
		openapigen.HandlerFromMux(apiHandler, api)
	})

	return &Server{
		router: r,
	}
}

func (s *Server) Start(addr string) error {
	s.service = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
	}

	return s.service.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.service == nil {
		return nil
	}
	return s.service.Shutdown(ctx)
}

func (s *Server) Routes() http.Handler {
	return s.router
}
