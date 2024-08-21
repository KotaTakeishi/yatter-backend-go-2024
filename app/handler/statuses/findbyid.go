package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	dto, err := h.statusUsecase.FindByID(ctx, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if dto == nil {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
