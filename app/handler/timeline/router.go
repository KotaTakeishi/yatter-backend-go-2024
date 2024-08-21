package timeline

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	timelineUsecase usecase.Timeline
}

// Create Handler for `/v1/timeline/`
func NewRouter(t usecase.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		timelineUsecase: t,
	}

	r.Get("/public", h.FindPublicTimelines)

	return r
}
