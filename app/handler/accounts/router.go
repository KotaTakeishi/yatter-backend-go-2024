package accounts

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	accountUsecase usecase.Account
}

// Create Handler for `/v1/accounts/`
func NewRouter(u usecase.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		accountUsecase: u,
	}
	r.Post("/", h.Create)

	return r
}
