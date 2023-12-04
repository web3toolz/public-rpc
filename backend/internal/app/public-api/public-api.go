package public_api

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"net"
	"net/http"
	"public-rpc/internal/adapters/cache"
	"public-rpc/internal/adapters/storage"
	"public-rpc/internal/app/public-api/query"
	"public-rpc/internal/config"
	publicapi "public-rpc/internal/ports/public-api"
)

type PublicAPIComponent struct {
	Cfg     config.PublicAPIConfig
	Logger  *zap.Logger
	Storage *storage.Storage
	Cache   cache.Cache
}

func (c *PublicAPIComponent) getHTTPHandler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(c.Cfg.Timeout))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/", func(r chi.Router) {
		r.Get("/", publicapi.GetRPCDataHandler(query.NewGetRPCDataHandler(c.Logger, c.Storage, c.Cache)))
	})

	return r
}

func (c *PublicAPIComponent) getHTTPServer(handler http.Handler) http.Server {
	return http.Server{
		Addr: net.JoinHostPort(c.Cfg.Host, c.Cfg.Port),

		Handler: handler,

		ReadHeaderTimeout: c.Cfg.Timeout,
		ReadTimeout:       c.Cfg.Timeout,
		WriteTimeout:      c.Cfg.Timeout,
		IdleTimeout:       c.Cfg.Timeout * 10,
	}
}

func (c *PublicAPIComponent) Run() error {
	handler := c.getHTTPHandler()
	server := c.getHTTPServer(handler)

	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
