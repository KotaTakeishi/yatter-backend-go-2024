package handler

import (
	"net/http"
	"time"

	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/accounts"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/health"
	"yatter-backend-go/app/handler/statuses"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(au usecase.Account, ar repository.Account) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(newCORS().Handler)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Mount("/v1/accounts", accounts.NewRouter(au))
	r.Mount("/v1/statuses", statuses.NewRouter(ar))
	r.Mount("/v1/health", health.NewRouter())
	r.Mount("/v1/auth", auth.NewRouter(ar))

	return r
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	})
}
