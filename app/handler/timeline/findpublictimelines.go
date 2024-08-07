package timeline

import (
	"encoding/json"
	"net/http"
)

func (h *handler) FindPublicTimelines(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, err := h.timelineUsecase.FindPublicTimelines(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
