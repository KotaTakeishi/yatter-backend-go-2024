package auth

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	ar repository.Account
}

func NewRouter(ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar}
	r.Post("/login", h.Login)

	return r
}
