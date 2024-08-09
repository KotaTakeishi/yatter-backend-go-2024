package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/auth"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	followee, err := h.accountUsecase.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, "failed to get followee account", http.StatusInternalServerError)
		return
	}
	if followee.Account == nil {
		http.Error(w, "followee account not found", http.StatusNotFound)
		return
	}

	follower := auth.AccountOf(ctx)

	dto, err := h.accountUsecase.Follow(ctx, follower.ID, followee.Account.ID)
	if err != nil {
		http.Error(w, "failed to follow account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Relation); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
