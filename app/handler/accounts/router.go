package accounts

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar             repository.Account
	accountUsecase usecase.Account
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account, u usecase.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		ar:             ar,
		accountUsecase: u,
	}

	r.Get("/{username}", h.FindByUsername)
	r.Post("/", h.Create)
	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		r.Put("/update_credentials", h.Update)
		r.Post("/{username}/follow", h.Follow)
	})

	return r
}
